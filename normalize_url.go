package main

import (
	"net/url"
	"strings"
)

func normalizeURL(origURL string) (string, error) {
	actualURL, err := url.Parse(origURL)
	if err != nil {
		return "", err
	}
	urlString := actualURL.Host + strings.TrimSuffix(actualURL.Path, "/")
	return urlString, nil
}
