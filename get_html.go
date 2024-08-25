package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// Make an HTTP GET request
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error getting website: %w", err)
	}
	defer res.Body.Close()

	// Check for HTTP error-level status codes
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP status code is an error-level code: %d", res.StatusCode)
	}

	// Check if the Content-Type is text/html
	contentType := res.Header.Get("Content-Type")
	if contentType != "text/html" && !strings.HasPrefix(contentType, "text/html;") {
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading website: %w", err)
	}

	return string(body), nil
}
