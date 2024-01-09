package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// defaultTimeout http clint request
var defaultTimeout int = 3 // in second

// Headers content headers parameter
type Headers map[string]string

// Add adds the key, value pair to the header.
func (h Headers) Add(key, value string) Headers {
	h[key] = value
	return h
}

// Get Header Value
func (h Headers) Get(key string) string {
	return h[key]
}

// RequestOptions provide option parameter client request
type RequestOptions struct {
	Payload       interface{}
	URL           string
	Header        Headers
	Method        string
	TimeoutSecond int
	Context       context.Context
}

// client setup new client
var client = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
		// ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		ForceAttemptHTTP2:     true,
	},
}

// ClientRequest http client request with option
func Request(option RequestOptions) (Response, error) {
	var rsp Response
	var err error

	start := time.Now()
	errTpl := "latency %s, %v, %s"

	// check option
	payload, err := castPayload(option.Payload)
	if err != nil {
		return rsp, err
	}

	if option.Context == nil {
		option.Context = context.Background()
	}

	// create request
	req, err := NewRequest(option.Context, option.Method, option.URL, payload)
	if err != nil {
		rsp.statusCode = http.StatusInternalServerError
		rsp.latency = time.Since(start)
		return rsp, fmt.Errorf(errTpl, rsp.Latency(), err.Error(), option.URL)
	}

	req.WithContext(option.Context)

	if option.TimeoutSecond < 1 {
		option.TimeoutSecond = defaultTimeout
	}

	// set headers
	for k, v := range option.Header {
		req.Header.Set(k, v)
	}

	client.Timeout = time.Duration(option.TimeoutSecond) * time.Second
	result, err := client.Do(req)
	rsp.latency = time.Since(start)

	// check is time out
	netErr, ok := err.(net.Error)
	if ok && netErr.Timeout() {
		rsp.statusCode = http.StatusRequestTimeout
	}

	// check err
	if err != nil {
		if !ok {
			rsp.statusCode = http.StatusInternalServerError
		}

		return rsp, fmt.Errorf(errTpl, rsp.Latency(), err.Error(), option.URL)
	}

	rsp.header = result.Header
	rsp.statusCode = result.StatusCode

	// read response body
	b, err := io.ReadAll(result.Body)
	if result != nil {
		result.Body.Close()
	}

	if err != nil {
		return rsp, fmt.Errorf(errTpl, rsp.Latency(), err.Error(), option.URL)
	}

	rsp.body = b

	return rsp, err
}

// castPayload to io.reader
func castPayload(param interface{}) (io.Reader, error) {
	var payload io.Reader
	var err error
	switch param.(type) {
	case string:
		payload = strings.NewReader(param.(string))
	case url.Values:
		payload = strings.NewReader(param.(url.Values).Encode())
	case []byte:
		payload = bytes.NewBuffer(param.([]byte))
	case io.Reader:
		payload = param.(io.Reader)
	case nil:
		payload = nil
	default:
		b, e := json.Marshal(param)
		if e == nil {
			payload = bytes.NewBuffer(b)
		}
		err = e
	}

	return payload, err
}
