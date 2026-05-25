package pagination

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const pagetotalHeader = "X-Pagination-Pagetotal"

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewPaginatedClient wraps base with a transport that copies the
// X-Pagination-Pagetotal response header into a synthetic `pagetotal` field
// on JSON-object responses, so that x-speakeasy-pagination (which only reads
// body fields) can drive auto-pagination on Cloudsmith list endpoints.
//
// If base is nil a default http.Client is used.
func NewPaginatedClient(base httpClient) *paginatedClient {
	if base == nil {
		base = &http.Client{}
	}
	return &paginatedClient{base: base}
}

type paginatedClient struct {
	base httpClient
}

func (c *paginatedClient) Do(req *http.Request) (*http.Response, error) {
	res, err := c.base.Do(req)
	if err != nil || res == nil || res.Body == nil {
		return res, err
	}
	return injectPagetotal(res), nil
}

func injectPagetotal(res *http.Response) *http.Response {
	raw := res.Header.Get(pagetotalHeader)
	if raw == "" {
		return res
	}
	total, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return res
	}
	if !strings.HasPrefix(res.Header.Get("Content-Type"), "application/json") {
		return res
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		res.Body = io.NopCloser(bytes.NewReader(body))
		return res
	}
	_ = res.Body.Close()
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(body, &obj); err != nil {
		res.Body = io.NopCloser(bytes.NewReader(body))
		return res
	}
	obj["pagetotal"] = json.RawMessage(strconv.FormatInt(total, 10))
	patched, err := json.Marshal(obj)
	if err != nil {
		res.Body = io.NopCloser(bytes.NewReader(body))
		return res
	}
	res.Body = io.NopCloser(bytes.NewReader(patched))
	res.ContentLength = int64(len(patched))
	res.Header.Set("Content-Length", strconv.Itoa(len(patched)))
	return res
}
