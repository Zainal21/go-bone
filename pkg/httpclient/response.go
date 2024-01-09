package httpclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"time"
)

var (
	errorCannotAddressable = errors.New("destination cannot addressable")
)

// ClientResponse representation of http client response
type Response struct {
	statusCode int
	body       []byte
	latency    time.Duration
	header     http.Header
}

// DecodeJSON decode response byte to struct
func (cr Response) DecodeJSON(dest interface{}) error {

	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errorCannotAddressable
	}

	return json.Unmarshal(cr.body, dest)
}

// String cast response byte to string
func (cr Response) String() string {
	return string(cr.body)
}

// RawByte return raw byte data
func (cr Response) RawByte() []byte {
	return cr.body
}

// Header return http header response
func (cr Response) Header() http.Header {
	return cr.header
}

// Header return http status response
func (cr Response) Status() int {
	return cr.statusCode
}

// Header return http status response
func (cr Response) Latency() time.Duration {
	return cr.latency
}
