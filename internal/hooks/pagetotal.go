package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const pagetotalHeader = "X-Pagination-Pagetotal"

// pagetotalHook copies the X-Pagination-Pagetotal response header into a
// synthetic `pagetotal` field on JSON-object responses, so that
// x-speakeasy-pagination (which only reads body fields) can drive
// auto-pagination on Cloudsmith list endpoints. No-ops when the header is
// absent, when the response is not JSON, or when the body is not an object.
type pagetotalHook struct{}

func (pagetotalHook) AfterSuccess(_ AfterSuccessContext, res *http.Response) (*http.Response, error) {
	if res == nil || res.Body == nil {
		return res, nil
	}
	raw := res.Header.Get(pagetotalHeader)
	if raw == "" {
		return res, nil
	}
	total, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return res, nil
	}
	if !strings.HasPrefix(res.Header.Get("Content-Type"), "application/json") {
		return res, nil
	}

	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return res, nil
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(body, &obj); err != nil {
		res.Body = io.NopCloser(bytes.NewReader(body))
		return res, nil
	}
	obj["pagetotal"] = json.RawMessage(strconv.FormatInt(total, 10))
	patched, err := json.Marshal(obj)
	if err != nil {
		res.Body = io.NopCloser(bytes.NewReader(body))
		return res, nil
	}
	res.Body = io.NopCloser(bytes.NewReader(patched))
	res.ContentLength = int64(len(patched))
	res.Header.Set("Content-Length", strconv.Itoa(len(patched)))
	return res, nil
}
