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
	url    string
	method method
	body   any
}

func FetchData[T any](payload Payload) (*T, error) {
	bodyBytes, err := json.Marshal(payload.body)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	req := &http.Request{
		RequestURI: payload.url,
		Method:     string(payload.method),
		Body:       io.NopCloser(bytes.NewBuffer(bodyBytes)),
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data T
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
