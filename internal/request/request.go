package request

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/JayantRautela/req/internal/model"
)


func SendRequest(config model.RequestConfig) (model.Response, error) {
	requestBody := bytes.NewBufferString(config.Body)

	req, err := http.NewRequest(
		config.Method,
		config.URL,
		requestBody,
	)

	if err != nil {
		return model.Response{}, err
	}

	for key, value := range config.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	startTime := time.Now()

	res, err := client.Do(req)

	if err != nil {
		return model.Response{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.Response{}, err
	}

	duration := time.Since(startTime)

	response := model.Response{
		StatusCode: res.StatusCode,
		Status:     res.Status,
		Headers:    res.Header,
		Body:       body,
		Duration:   duration,
	}

	return response, nil
}