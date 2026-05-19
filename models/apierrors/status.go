package apierrors

import (
	"errors"
	"net/http"
)

// IsNotFound reports whether err wraps an APIError with HTTP 404.
func IsNotFound(err error) bool {
	return statusCode(err) == http.StatusNotFound
}

// IsConflict reports whether err wraps an APIError with HTTP 409.
func IsConflict(err error) bool {
	return statusCode(err) == http.StatusConflict
}

// IsUnauthorized reports whether err wraps an APIError with HTTP 401.
func IsUnauthorized(err error) bool {
	return statusCode(err) == http.StatusUnauthorized
}

// IsForbidden reports whether err wraps an APIError with HTTP 403.
func IsForbidden(err error) bool {
	return statusCode(err) == http.StatusForbidden
}

func statusCode(err error) int {
	var ae *APIError
	if errors.As(err, &ae) {
		return ae.StatusCode
	}
	return 0
}
