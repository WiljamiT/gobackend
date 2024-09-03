package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from the headers of an HTTP req
// Example:
// Authorization: ApiKey {insert apikey here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth header")
	}

	return vals[1], nil
}
