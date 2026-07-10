package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func Fetch[T any](ctx context.Context, payload Payload) (*T, error) {
	req, err := formatRequest(ctx, payload)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var body any
		json.NewDecoder(res.Body).Decode(&body)

		return nil, fmt.Errorf("Something went wrong, please try later")
	}

	var data T
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func formatRequest(ctx context.Context, payload Payload) (*http.Request, error) {
	var body io.Reader
	if payload.Body != nil {
		bodyBytes, err := json.Marshal(payload.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, string(payload.Method), payload.Url, body)
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
