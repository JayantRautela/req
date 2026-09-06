package normalizer

import (
	"strings"

	"github.com/JayantRautela/req/internal/model"
)

func NormalizeRequest(config model.RequestConfig) model.RequestConfig {
	config.URL = strings.TrimSpace(config.URL)

	config.Method = strings.ToUpper(
		strings.TrimSpace(config.Method),
	)

	config.Body = strings.TrimSpace(config.Body)

	config.Headers = normalizeHeaders(config.Headers)

	return config
}

func normalizeHeaders(headers map[string]string) map[string]string {
	normalizedHeaders := make(map[string]string, len(headers))

	for key, value := range headers {
		normalizedKey := strings.TrimSpace(key)
		normalizedValue := strings.TrimSpace(value)

		normalizedHeaders[normalizedKey] = normalizedValue
	}

	return normalizedHeaders
}