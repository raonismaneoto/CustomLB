package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpResponse struct {
	Body       []byte
	Headers    http.Header
	StatusCode int
}

var (
	Client HTTPClient = &http.Client{}
)

type Request struct {
	Body     interface{}
	Endpoint string
	Method   string
	Headers  http.Header
}

func SendRequest(req *Request) (*HttpResponse, error) {
	if req.Body == nil {
		req.Body = &map[string]interface{}{}
	}

	parsedBody, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(req.Method, req.Endpoint, bytes.NewBuffer(parsedBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header = req.Headers

	resp, err := Client.Do(httpReq)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &HttpResponse{Body: respBody, Headers: resp.Header, StatusCode: resp.StatusCode}, nil
}
