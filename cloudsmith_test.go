package cloudsmith_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type headerRecorder struct {
	header http.Header
}

func (r *headerRecorder) Do(req *http.Request) (*http.Response, error) {
	r.header = req.Header.Clone()
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{}`)),
		Request:    req,
	}, nil
}

type urlRecorder struct {
	requestURL string
}

func (r *urlRecorder) Do(req *http.Request) (*http.Response, error) {
	r.requestURL = req.URL.String()
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{}`)),
		Request:    req,
	}, nil
}

type pagetotalInnerClient struct{ called bool }

func (c *pagetotalInnerClient) Do(req *http.Request) (*http.Response, error) {
	c.called = true
	return &http.Response{
		StatusCode: http.StatusOK,
		Header: http.Header{
			"Content-Type":          []string{"application/json"},
			"X-Pagination-Pagetotal": []string{"3"},
		},
		Body:    io.NopCloser(strings.NewReader(`{}`)),
		Request: req,
	}, nil
}

func minimalMetadataListRequest() operations.MetadataPackagesListRequest {
	return operations.MetadataPackagesListRequest{PackageSlugPerm: "test-pkg"}
}

func TestNew_SetsAPIKeyFromEnv(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_KEY", "env-key")
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com")

	recorder := &headerRecorder{}
	sdk := cloudsmith.New(cloudsmith.WithClient(recorder))
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)
	require.NotNil(t, recorder.header)

	assert.Equal(t, "env-key", recorder.header.Get("X-Api-Key"))
}

func TestNew_ExplicitSecurityOverridesEnv(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_KEY", "env-key")
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com")

	recorder := &headerRecorder{}
	sdk := cloudsmith.New(
		cloudsmith.WithSecurity("explicit-key"),
		cloudsmith.WithClient(recorder),
	)
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)
	require.NotNil(t, recorder.header)

	assert.Equal(t, "explicit-key", recorder.header.Get("X-Api-Key"))
}

func TestNew_NoCredentials(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_KEY", "")
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com")

	recorder := &headerRecorder{}
	sdk := cloudsmith.New(cloudsmith.WithClient(recorder))
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)
	require.NotNil(t, recorder.header)

	assert.Empty(t, recorder.header.Get("X-Api-Key"))
}

func TestNew_PaginatedClientWrapsPagetotal(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_KEY", "key")
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com")

	inner := &pagetotalInnerClient{}
	sdk := cloudsmith.New(cloudsmith.WithClient(inner))
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)

	assert.True(t, inner.called)
}

func TestNew_SetsServerURLFromEnv(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com")
	t.Setenv("CLOUDSMITH_API_KEY", "test-key")

	recorder := &urlRecorder{}
	sdk := cloudsmith.New(cloudsmith.WithClient(recorder))
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)

	assert.Contains(t, recorder.requestURL, "api.example.com/v2/")
}

func TestNew_SetsServerURLFromEnv_AlreadyHasV2Path(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com/v2/")
	t.Setenv("CLOUDSMITH_API_KEY", "test-key")

	recorder := &urlRecorder{}
	sdk := cloudsmith.New(cloudsmith.WithClient(recorder))
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)

	assert.Contains(t, recorder.requestURL, "api.example.com/v2/")
	assert.NotContains(t, recorder.requestURL, "/v2/v2/")
}

func TestNew_SetsServerURLFromEnv_LegacyV1PathReplacedWithV2(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "https://api.example.com/v1")
	t.Setenv("CLOUDSMITH_API_KEY", "test-key")

	recorder := &urlRecorder{}
	sdk := cloudsmith.New(cloudsmith.WithClient(recorder))
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)

	assert.Contains(t, recorder.requestURL, "api.example.com/v2/")
	assert.NotContains(t, recorder.requestURL, "/v1")
}

func TestNew_NormalizesExplicitServerURL(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "")
	t.Setenv("CLOUDSMITH_API_KEY", "test-key")

	recorder := &urlRecorder{}
	sdk := cloudsmith.New(
		cloudsmith.WithServerURL("https://explicit.example.com"),
		cloudsmith.WithClient(recorder),
	)
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)

	assert.Contains(t, recorder.requestURL, "explicit.example.com/v2/")
}

func TestNew_NormalizesExplicitServerURL_LegacyV1PathReplacedWithV2(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "")
	t.Setenv("CLOUDSMITH_API_KEY", "test-key")

	recorder := &urlRecorder{}
	sdk := cloudsmith.New(
		cloudsmith.WithServerURL("https://explicit.example.com/v1"),
		cloudsmith.WithClient(recorder),
	)
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)

	assert.Contains(t, recorder.requestURL, "explicit.example.com/v2/")
	assert.NotContains(t, recorder.requestURL, "/v1")
}

func TestNew_ExplicitServerURLTakesPrecedenceOverEnv(t *testing.T) {
	t.Setenv("CLOUDSMITH_API_HOST", "https://env.example.com")
	t.Setenv("CLOUDSMITH_API_KEY", "test-key")

	recorder := &urlRecorder{}
	sdk := cloudsmith.New(
		cloudsmith.WithServerURL("https://explicit.example.com"),
		cloudsmith.WithClient(recorder),
	)
	_, err := sdk.Metadata.MetadataPackagesList(context.Background(), minimalMetadataListRequest())
	require.NoError(t, err)

	assert.Contains(t, recorder.requestURL, "explicit.example.com/v2/")
	assert.NotContains(t, recorder.requestURL, "env.example.com")
}
