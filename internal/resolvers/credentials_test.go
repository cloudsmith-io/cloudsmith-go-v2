package resolvers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/cloudsmith-io/cloudsmith-go-v2/internal/resolvers"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubProvider struct {
	cred components.Security
	err  error
}

func (s stubProvider) Resolve(_ context.Context) (components.Security, error) {
	return s.cred, s.err
}

func TestCredentialProviderChain_EmptyChain(t *testing.T) {
	chain := resolvers.NewCredentialProviderChain()
	_, err := chain.Resolve(context.Background())
	assert.ErrorIs(t, err, resolvers.ErrNoCredentials)
}

func TestCredentialProviderChain_FirstProviderSucceeds(t *testing.T) {
	want := components.Security{Apikey: "first-key"}
	chain := resolvers.NewCredentialProviderChain(
		stubProvider{cred: want},
		stubProvider{cred: components.Security{Apikey: "second-key"}},
	)
	got, err := chain.Resolve(context.Background())
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestCredentialProviderChain_SkipsNoCredentials(t *testing.T) {
	want := components.Security{Apikey: "second-key"}
	chain := resolvers.NewCredentialProviderChain(
		stubProvider{err: resolvers.ErrNoCredentials},
		stubProvider{cred: want},
	)
	got, err := chain.Resolve(context.Background())
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestCredentialProviderChain_PropagatesHardError(t *testing.T) {
	hard := errors.New("keyring locked")
	chain := resolvers.NewCredentialProviderChain(
		stubProvider{err: hard},
		stubProvider{cred: components.Security{Apikey: "fallback"}},
	)
	_, err := chain.Resolve(context.Background())
	assert.ErrorIs(t, err, hard)
}

func TestCredentialProviderChain_AllProvidersMiss(t *testing.T) {
	chain := resolvers.NewCredentialProviderChain(
		stubProvider{err: resolvers.ErrNoCredentials},
		stubProvider{err: resolvers.ErrNoCredentials},
	)
	_, err := chain.Resolve(context.Background())
	assert.ErrorIs(t, err, resolvers.ErrNoCredentials)
}

func TestEnvProvider_KeySet(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_KEY", "test-api-key")
	cred, err := resolvers.EnvProvider{}.Resolve(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "test-api-key", cred.Apikey)
}

func TestEnvProvider_KeyEmpty(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_KEY", "")
	_, err := resolvers.EnvProvider{}.Resolve(context.Background())
	assert.ErrorIs(t, err, resolvers.ErrNoCredentials)
}

func TestEnvProvider_KeyUnset(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_KEY", "")
	_, err := resolvers.EnvProvider{}.Resolve(context.Background())
	assert.ErrorIs(t, err, resolvers.ErrNoCredentials)
}
