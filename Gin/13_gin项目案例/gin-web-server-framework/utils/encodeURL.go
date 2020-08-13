package utils

import (
	"net/url"
)

// EncodeURL add and encode parameters.
func EncodeURL(api string, params url.Values) (string, error) {
	url, err := url.Parse(api)
	if err != nil {
		return "", err
	}

	query := url.Query()

	for k := range params {
		query.Set(k, params.Get(k))
	}

	url.RawQuery = query.Encode()

	return url.String(), nil
}
