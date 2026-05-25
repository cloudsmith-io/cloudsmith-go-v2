package pagination_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudsmith-io/cloudsmith-go-v2/internal/pagination"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPaginatedClient_InjectsPagetotal(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Pagination-Pagetotal", "42")
		_, _ = w.Write([]byte(`{"results":[]}`))
	}))
	defer srv.Close()

	client := pagination.NewPaginatedClient(srv.Client())
	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res, err := client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var obj map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(body, &obj))
	assert.Equal(t, json.RawMessage("42"), obj["pagetotal"])
}

func TestNewPaginatedClient_NoHeaderPassesThrough(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"results":[]}`))
	}))
	defer srv.Close()

	client := pagination.NewPaginatedClient(srv.Client())
	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res, err := client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var obj map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(body, &obj))
	_, hasPagetotal := obj["pagetotal"]
	assert.False(t, hasPagetotal)
}

func TestNewPaginatedClient_NonJSONPassesThrough(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("X-Pagination-Pagetotal", "5")
		_, _ = w.Write([]byte("hello"))
	}))
	defer srv.Close()

	client := pagination.NewPaginatedClient(srv.Client())
	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res, err := client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, "hello", string(body))
}

func TestNewPaginatedClient_InvalidHeaderPassesThrough(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Pagination-Pagetotal", "not-a-number")
		_, _ = w.Write([]byte(`{"results":[]}`))
	}))
	defer srv.Close()

	client := pagination.NewPaginatedClient(srv.Client())
	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res, err := client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var obj map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(body, &obj))
	_, hasPagetotal := obj["pagetotal"]
	assert.False(t, hasPagetotal)
}
