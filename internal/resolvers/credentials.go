package resolvers

import (
	"context"
	"errors"
	"os"

	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

// ErrNoCredentials is returned by a CredentialProvider when it has no
// credentials to offer. The chain treats this as a soft miss and tries the
// next provider; any other error stops the chain immediately.
var ErrNoCredentials = errors.New("no credentials available")

// CredentialProvider resolves Cloudsmith API credentials from a single source.
type CredentialProvider interface {
	Resolve(ctx context.Context) (components.Security, error)
}

// CredentialProviderChain resolves credentials by trying each provider in
// order and returning the first successful result.
type CredentialProviderChain struct {
	providers []CredentialProvider
}

// NewCredentialProviderChain constructs a chain from the given providers.
func NewCredentialProviderChain(providers ...CredentialProvider) CredentialProviderChain {
	return CredentialProviderChain{providers: providers}
}

// Resolve iterates providers. ErrNoCredentials causes the chain to advance to
// the next provider; any other error is returned immediately.
func (c CredentialProviderChain) Resolve(ctx context.Context) (components.Security, error) {
	for _, p := range c.providers {
		cred, err := p.Resolve(ctx)
		if errors.Is(err, ErrNoCredentials) {
			continue
		}
		if err != nil {
			return components.Security{}, err
		}
		return cred, nil
	}
	return components.Security{}, ErrNoCredentials
}

// EnvProvider resolves credentials from the CLOUDSMITH_API_KEY environment variable.
type EnvProvider struct{}

// Resolve returns ErrNoCredentials if CLOUDSMITH_API_KEY is unset or empty.
func (EnvProvider) Resolve(_ context.Context) (components.Security, error) {
	key := os.Getenv("CLOUDSMITH_API_KEY")
	if key == "" {
		return components.Security{}, ErrNoCredentials
	}
	return components.Security{Apikey: key}, nil
}

// DefaultCredentialProviderChain returns the chain used when no explicit
// security is configured. Providers are tried in this order:
//
//  1. CLOUDSMITH_API_KEY environment variable
func DefaultCredentialProviderChain() CredentialProviderChain {
	return NewCredentialProviderChain(EnvProvider{})
}
