package validator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/JayantRautela/req/internal/model"
)

func ValidateRequest(config model.RequestConfig) error {
	if err := validateURL(config.URL); err != nil {
		return err
	}

	if err := validateMethod(config.Method); err != nil {
		return err
	}

	if err := validateBody(config.Body); err != nil {
		return err
	}

	return nil
}

func validateURL(rawURL string) error {
	if rawURL == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("URL must use http or https")
	}

	if parsedURL.Host == "" {
		return fmt.Errorf("URL must contain a host")
	}

	return nil
}

func validateMethod(method string) error {
	switch method {
	case http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions:
		return nil

	default:
		return fmt.Errorf("unsupported HTTP method: %s", method)
	}
}

func validateBody(body string) error {
	if body == "" {
		return nil
	}

	if !json.Valid([]byte(body)) {
		return fmt.Errorf("request body contains invalid JSON")
	}

	return nil
}
