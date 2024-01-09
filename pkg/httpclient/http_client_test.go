// Package utils tests
// @author Daud Valentino
package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestClientRequest(t *testing.T) {

	t.Parallel()
	type response struct {
		Code int `json:"code"`
	}

	type payload struct {
		Code      int  `json:"code"`
		WantError bool `json:"want_error"`
	}
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var p payload

		_ = json.NewDecoder(r.Body).Decode(&p)

		w.Header().Set("Content-Type", "application/json")

		var rsp []byte
		switch p.Code {
		case 1000:
			rsp = []byte(`{"code":1000}`)
		case 2000:
			rsp = []byte(`{"code":2000}`)
		case 2005:
			rsp = []byte(`{"code":2005}`)
		case 7000:
			time.Sleep(3 * time.Second)
			rsp = []byte(`{"code":1000}`)
		case 500:
			w.WriteHeader(http.StatusInternalServerError)
			rsp = []byte(`{"code":500}`)
		case 999:
			w.Header().Set("Content-Length", "1")
		case 998:
			w.Header().Set("Content-Length", "1")
			rsp = []byte(`{"code":998}`)
		default:
			rsp = nil
		}

		w.Write(rsp)

	}))

	defer fakeServer.Close()

	vl := url.Values{}
	vl.Add("number", "1000")

	testCases := []struct {
		WantCode  int
		WantError bool
		payload   interface{}
		TimeOut   int
		Url       string
		Method    string
		Ctx       context.Context
		Headers   map[string]string
		Name      string
	}{
		{
			Name:      "scenario #1 error cast payload",
			WantCode:  1000,
			WantError: true,
			TimeOut:   2,
			payload:   func() {},
			Url:       "",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},

		{
			Name:      "scenario #2 error create new http request",
			WantCode:  0,
			WantError: true,
			TimeOut:   2,
			payload:   nil,
			Url:       fakeServer.URL,
			Method:    "/",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			Name:      "scenario #3 http timeout",
			WantCode:  0,
			WantError: true,
			TimeOut:   1,
			payload:   []byte(`{"code": 7000}`),
			Url:       fakeServer.URL,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},

		{
			Name:      "scenario #4 http unsupport protocol scheme",
			WantCode:  0,
			WantError: true,
			TimeOut:   0,
			payload:   `{"code": 1000}`,
			Url:       "fakeServer.URL",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},

		{
			Name:      "scenario #5 success",
			WantCode:  1000,
			WantError: false,
			TimeOut:   2,
			payload:   []byte(`{"code": 1000}`),
			Url:       fakeServer.URL,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},

		{
			Name:      "scenario #6 wrong param with header",
			WantCode:  0,
			WantError: false,
			TimeOut:   2,
			payload:   url.Values{},
			Url:       fakeServer.URL,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},

		{
			Name:      "scenario #7 success with payload io reader",
			WantCode:  1000,
			WantError: false,
			TimeOut:   2,
			payload:   strings.NewReader(`{"code": 1000}`),
			Url:       fakeServer.URL,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},

		{
			Name:      "scenario #8 payload set nil",
			WantCode:  0,
			WantError: false,
			TimeOut:   2,
			payload:   nil,
			Url:       fakeServer.URL,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},

		{
			Name:      "scenario #9 success with payload struct",
			WantCode:  1000,
			WantError: false,
			TimeOut:   2,
			payload: struct {
				Code int64 `json:"code"`
			}{
				Code: 1000,
			},
			Url: fakeServer.URL,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},
		{
			Name:      "scenario #10 ioutil.ReadAll error",
			WantCode:  0,
			WantError: false,
			TimeOut:   2,
			payload: struct {
				Code int64 `json:"code"`
			}{
				Code: 999,
			},
			Url: fakeServer.URL,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},
		{
			Name:      "scenario #11 client do error",
			WantCode:  0,
			WantError: false,
			TimeOut:   2,
			payload:   "test 123",
			Url:       fakeServer.URL,
			Headers: map[string]string{
				"Content-Type": "application/json",
				"Location":     "jakrta@gmail.com||",
			},
			Ctx:    context.Background(),
			Method: http.MethodPost,
		},
	}

	for i, x := range testCases {
		result, err := Request(RequestOptions{
			Payload:       x.payload,
			URL:           x.Url,
			TimeoutSecond: x.TimeOut,
			Method:        x.Method,
			Context:       x.Ctx,
			Header:        x.Headers,
		})

		rsp := response{}
		_ = json.Unmarshal(result.RawByte(), &rsp)

		if x.WantError {
			if err == nil {
				t.Fatalf("%s expected error not nil, got: %v, index: %v result %s", x.Name, err, i, result.String())
				continue
			}

			t.Logf("%s expected error not nil, got: %v, index: %v result %s", x.Name, err, i, result.String())
			continue
		}

		if !x.WantError {
			if x.WantCode != rsp.Code {
				t.Fatalf("%s expected  %d, got: %d, index: %v result %s", x.Name, x.WantCode, rsp.Code, i, result.String())
				continue
			}

			t.Logf("%s expected  %d, got: %d, index: %v result %s", x.Name, x.WantCode, rsp.Code, i, result.String())
		}
	}

}

func TestHeaders_Add(t *testing.T) {
	t.Parallel()
	h := Headers{}
	h.Add("scenarion#1", "1")
	t.Run("scenario add value", func(t *testing.T) {
		r := h.Get("scenarion#1")

		if r == "1" {
			t.Logf("expected %s, actual %s", "1", r)
		} else {
			t.Fatalf("expected %s, actual %s", "1", r)
		}
	})
}

func TestHeaders_Get(t *testing.T) {
	t.Parallel()
	h := Headers{}
	h.Add("scenario#2", "2")
	t.Run("scenario get value", func(t *testing.T) {
		r := h.Get("scenario#2")

		if r == "2" {
			t.Logf("expected %s, actual %s", "2", r)
		} else {
			t.Fatalf("expected %s, actual %s", "2", r)
		}
	})
}
