package resolvers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/cloudsmith-io/cloudsmith-go-v2/internal/resolvers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubEndpointProvider struct {
	url string
	err error
}

func (s stubEndpointProvider) Resolve(_ context.Context) (string, error) {
	return s.url, s.err
}

func TestEndpointProviderChain_EmptyChain(t *testing.T) {
	chain := resolvers.NewEndpointProviderChain()
	_, err := chain.Resolve(context.Background())
	assert.ErrorIs(t, err, resolvers.ErrNoEndpoint)
}

func TestEndpointProviderChain_FirstProviderSucceeds(t *testing.T) {
	chain := resolvers.NewEndpointProviderChain(
		stubEndpointProvider{url: "https://first.example.com"},
		stubEndpointProvider{url: "https://second.example.com"},
	)
	got, err := chain.Resolve(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "https://first.example.com", got)
}

func TestEndpointProviderChain_SkipsNoEndpoint(t *testing.T) {
	chain := resolvers.NewEndpointProviderChain(
		stubEndpointProvider{err: resolvers.ErrNoEndpoint},
		stubEndpointProvider{url: "https://second.example.com"},
	)
	got, err := chain.Resolve(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "https://second.example.com", got)
}

func TestEndpointProviderChain_PropagatesHardError(t *testing.T) {
	hard := errors.New("config file parse error")
	chain := resolvers.NewEndpointProviderChain(
		stubEndpointProvider{err: hard},
		stubEndpointProvider{url: "https://fallback.example.com"},
	)
	_, err := chain.Resolve(context.Background())
	assert.ErrorIs(t, err, hard)
}

func TestEndpointProviderChain_AllProvidersMiss(t *testing.T) {
	chain := resolvers.NewEndpointProviderChain(
		stubEndpointProvider{err: resolvers.ErrNoEndpoint},
		stubEndpointProvider{err: resolvers.ErrNoEndpoint},
	)
	_, err := chain.Resolve(context.Background())
	assert.ErrorIs(t, err, resolvers.ErrNoEndpoint)
}

func TestEnvEndpointProvider_HostSet(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com")
	url, err := resolvers.EnvEndpointProvider{}.Resolve(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "https://api.example.com", url)
}

func TestEnvEndpointProvider_LegacyV1PathStripped(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com/v1")
	url, err := resolvers.EnvEndpointProvider{}.Resolve(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "https://api.example.com", url)
}

func TestEnvEndpointProvider_HostEmpty(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "")
	_, err := resolvers.EnvEndpointProvider{}.Resolve(context.Background())
	assert.ErrorIs(t, err, resolvers.ErrNoEndpoint)
}

func TestEnvEndpointProvider_HostUnset(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "")
	_, err := resolvers.EnvEndpointProvider{}.Resolve(context.Background())
	assert.ErrorIs(t, err, resolvers.ErrNoEndpoint)
}
