package httpclient

import (
	"net/http"
	"testing"
	"time"
)

func TestResponse_DecodeJSON(t *testing.T) {
	t.Parallel()
	rsp := Response{
		statusCode: 0,
		body:       []byte(`{"code": 1000}`),
		latency:    0,
		header:     nil,
	}

	t.Run("scenario decode success", func(t *testing.T) {

		result := struct {
			Code int64 `json:"code"`
		}{}

		expected := int64(1000)
		rsp.DecodeJSON(&result)

		if result.Code == expected {
			t.Logf("expected %d, actual %d", expected, result.Code)
		} else {
			t.Fatalf("expected %d, actual %d", expected, result.Code)
		}
	})

	t.Run("scenario decode cannot addressable", func(t *testing.T) {

		result := struct {
			Code int64 `json:"code"`
		}{}

		expected := errorCannotAddressable
		err := rsp.DecodeJSON(result)

		if expected == err {
			t.Logf("expected %v, actual %v", expected, err)
		} else {
			t.Fatalf("expected %v, actual %v", expected, err)
		}
	})
}

func TestResponse_RawByte(t *testing.T) {
	t.Parallel()
	rsp := Response{
		statusCode: 0,
		body:       []byte(`{"code": 1000}`),
		latency:    0,
		header:     nil,
	}

	t.Run("scenario get value", func(t *testing.T) {

		expected := []byte(`{"code": 1000}`)
		result := rsp.RawByte()

		if len(result) == len(expected) {
			t.Logf("expected %d, actual %d", len(expected), len(result))
		} else {
			t.Fatalf("expected %d, actual %d", len(expected), len(result))
		}
	})
}

func TestResponse_String(t *testing.T) {
	t.Parallel()
	rsp := Response{
		statusCode: 0,
		body:       []byte(`{"code": 1000}`),
		latency:    0,
		header:     nil,
	}

	t.Run("scenario get value", func(t *testing.T) {

		expected := `{"code": 1000}`
		result := rsp.String()

		if expected == result {
			t.Logf("expected %s, actual %s", expected, result)
		} else {
			t.Fatalf("expected %s, actual %s", expected, result)
		}
	})
}

func TestResponse_Status(t *testing.T) {
	t.Parallel()
	rsp := Response{
		statusCode: http.StatusOK,
		body:       nil,
		latency:    0,
		header:     nil,
	}

	t.Run("scenario get value", func(t *testing.T) {

		expected := http.StatusOK
		result := rsp.Status()

		if expected == result {
			t.Logf("expected %d, actual %d", expected, result)
		} else {
			t.Fatalf("expected %d actual %d", expected, result)
		}
	})
}

func TestResponse_Latency(t *testing.T) {
	t.Parallel()
	rsp := Response{
		statusCode: http.StatusOK,
		body:       nil,
		latency:    0,
		header:     nil,
	}

	t.Run("scenario get latency", func(t *testing.T) {

		expected := time.Duration(0)
		result := rsp.Latency()

		if expected == result {
			t.Logf("expected %d, actual %d", expected, result)
		} else {
			t.Fatalf("expected %d actual %d", expected, result)
		}
	})
}

func TestResponse_Header(t *testing.T) {
	t.Parallel()

	expected := http.Header{}
	expected.Set("#1", "1")
	rsp := Response{
		statusCode: http.StatusOK,
		body:       nil,
		latency:    0,
		header:     expected,
	}

	t.Run("scenario get header", func(t *testing.T) {

		actual := rsp.Header()

		if expected.Get("#1") == actual.Get("#1") {
			t.Logf("expected %s, actual %s", expected.Get("#1"), actual.Get("#1"))
		} else {
			t.Fatalf("expected %s actual %s", expected.Get("#1"), actual.Get("#1"))
		}
	})
}
