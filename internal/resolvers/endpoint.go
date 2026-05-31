package resolvers

import (
	"context"
	"errors"
	"os"
	"strings"
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

// StaticEndpointProvider resolves a fixed, explicitly-configured API host (for
// example one supplied via WithServerURL). It returns ErrNoEndpoint when the
// configured value is empty so the chain falls through to lower-precedence
// providers.
type StaticEndpointProvider struct {
	URL string
}

// Resolve returns ErrNoEndpoint if URL is empty.
func (p StaticEndpointProvider) Resolve(_ context.Context) (string, error) {
	if p.URL == "" {
		return "", ErrNoEndpoint
	}
	return p.URL, nil
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

// NormalizingEndpointProvider wraps an EndpointProviderChain and normalizes the
// resolved host: a trailing slash and legacy /v1 suffix are stripped, the path
// is anchored at /v2, and a trailing slash is restored.
type NormalizingEndpointProvider struct {
	chain EndpointProviderChain
}

// Resolve resolves the underlying chain and normalizes the resulting host. Any
// error from the chain (including ErrNoEndpoint) is propagated unchanged.
func (p NormalizingEndpointProvider) Resolve(ctx context.Context) (string, error) {
	host, err := p.chain.Resolve(ctx)
	if err != nil {
		return "", err
	}
	return normalizeServerURL(host), nil
}

// normalizeServerURL strips a trailing slash and legacy /v1 suffix from a
// resolved host, ensures the path is anchored at /v2, and restores the trailing
// slash expected by the SDK.
func normalizeServerURL(host string) string {
	host = strings.TrimSuffix(strings.TrimRight(host, "/"), "/v1")
	if !strings.HasSuffix(host, "/v2") {
		host += "/v2"
	}
	return host + "/"
}

// DefaultEndpointProviderChain returns the resolver used to determine the API
// host. Any explicitURL values (typically the server URL set via
// WithServerURL) take highest precedence, followed by:
//
//  1. CLOUDSMITH_API_HOST environment variable
//
// The resolved host is normalized via NormalizingEndpointProvider.
func DefaultEndpointProviderChain(explicitURL ...string) NormalizingEndpointProvider {
	providers := make([]EndpointProvider, 0, len(explicitURL)+1)
	for _, url := range explicitURL {
		providers = append(providers, StaticEndpointProvider{URL: url})
	}
	providers = append(providers, EnvEndpointProvider{})
	return NormalizingEndpointProvider{chain: NewEndpointProviderChain(providers...)}
}
