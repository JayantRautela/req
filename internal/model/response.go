package model

import "time"

type Response struct {
	StatusCode int
	Status string
	Headers map[string][]string
	Duration time.Duration
	Body []byte
}