package resolvers

import (
	"context"
	"errors"
	"os"
)

const envAPIHost = "CLOUDSMITH_API_HOST"

// ErrNoEndpoint is returned by an EndpointProvider when it has no endpoint to
// offer. The chain treats this as a soft miss and tries the next provider; any
// other error stops the chain immediately.
var ErrNoEndpoint = errors.New("no endpoint available")

// EndpointProvider resolves a Cloudsmith API host URL from a single source.
type EndpointProvider interface {
	Resolve(ctx context.Context) (string, error)
}

// EndpointProviderChain resolves an endpoint by trying each provider in order
// and returning the first successful result.
type EndpointProviderChain struct {
	providers []EndpointProvider
}

// NewEndpointProviderChain constructs a chain from the given providers.
func NewEndpointProviderChain(providers ...EndpointProvider) EndpointProviderChain {
	return EndpointProviderChain{providers: providers}
}

// Resolve iterates providers. ErrNoEndpoint causes the chain to advance to the
// next provider; any other error is returned immediately.
func (c EndpointProviderChain) Resolve(ctx context.Context) (string, error) {
	for _, p := range c.providers {
		url, err := p.Resolve(ctx)
		if errors.Is(err, ErrNoEndpoint) {
			continue
		}
		if err != nil {
			return "", err
		}
		return url, nil
	}
	return "", ErrNoEndpoint
}

// EnvEndpointProvider resolves the API host from the CLOUDSMITH_API_HOST
// environment variable.
type EnvEndpointProvider struct{}

// Resolve returns ErrNoEndpoint if CLOUDSMITH_API_HOST is unset or empty.
func (EnvEndpointProvider) Resolve(_ context.Context) (string, error) {
	host := os.Getenv(envAPIHost)
	if host == "" {
		return "", ErrNoEndpoint
	}
	return host, nil
}

// DefaultEndpointProviderChain returns the chain used when no explicit server
// URL is configured. Providers are tried in this order:
//
//  1. CLOUDSMITH_API_HOST environment variable
func DefaultEndpointProviderChain() EndpointProviderChain {
	return NewEndpointProviderChain(EnvEndpointProvider{})
}
