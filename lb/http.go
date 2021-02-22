package lb

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
	Client       HTTPClient                                        = &http.Client{}
)

func SendRequest(endpoint string, incomingReq *Request) (*HttpResponse, error) {
	if incomingReq.Body == nil {
		incomingReq.Body = &map[string]string{}
	}

	parsedBody, err := json.Marshal(incomingReq.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(incomingReq.Method, endpoint + incomingReq.ResourcePath, bytes.NewBuffer(parsedBody))
	if err != nil {
		return nil, err
	}

	req.Header = incomingReq.Headers

	resp, err := Client.Do(req)

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
