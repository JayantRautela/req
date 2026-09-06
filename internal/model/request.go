package model

type RequestConfig struct {
	URL     string
	Method string
	Body    string
	Headers map[string]string
}
