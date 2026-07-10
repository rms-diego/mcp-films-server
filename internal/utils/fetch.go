package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type method string

const (
	GET    method = "GET"
	POST   method = "POST"
	PUT    method = "PUT"
	PATCH  method = "PATCH"
	DELETE method = "DELETE"
)

type Payload struct {
	Url     string
	Method  method
	Body    *any
	Headers map[string]string
}

func Fetch[T any](payload Payload) (*T, error) {
	req, err := formatRequest(payload)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data T
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func formatRequest(payload Payload) (*http.Request, error) {
	bodyBytes, err := json.Marshal(payload.Body)
	if err != nil {
		return nil, err
	}

	var body io.Reader
	if bodyBytes != nil {
		body = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(string(payload.Method), payload.Url, body)
	if err != nil {
		return nil, err
	}

	if len(payload.Headers) > 0 {
		for key, value := range payload.Headers {
			req.Header.Set(key, value)
		}
	}

	return req, nil
}
