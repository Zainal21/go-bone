// Package appctx
package appctx

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	rsp    *Response
	oneRsp sync.Once
)

// Response presentation contract object
type Response struct {
	Code      int         `json:"-"`
	Status    bool        `json:"status"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Entity    string      `json:"entity,omitempty"`
	State     string      `json:"state,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
	Message   interface{} `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	lang      string      `json:"-"`
	msgKey    string
}

// WithCode setter response var name
func (r *Response) WithCode(c int) *Response {
	r.Status = true

	if c > fiber.StatusCreated {
		r.Status = false
	}
	r.Code = c
	return r
}

func (r *Response) WithError(v interface{}) *Response {
	r.Errors = v
	return r
}

// WithEntity setter entity response
func (r *Response) WithEntity(e string) *Response {
	r.Entity = e
	return r
}

// WithState setter state response
func (r *Response) WithState(s string) *Response {
	r.State = s
	return r
}

// WithData setter data response
func (r *Response) WithData(v interface{}) *Response {
	r.Data = v
	return r
}

func (r *Response) WithMsgKey(v string) *Response {
	r.msgKey = v
	return r
}

// WithMeta setter meta data response
func (r *Response) WithMeta(v interface{}) *Response {
	r.Meta = v
	return r
}

// WithLang setter language response
func (r *Response) WithLang(v string) *Response {
	r.lang = v
	return r
}

// WithMessage setter custom message response
func (r *Response) WithMessage(v interface{}) *Response {
	if v != nil {
		r.Message = v
	}

	return r
}

func (r *Response) Byte() []byte {
	b, _ := json.Marshal(r)
	return b
}

// NewResponse initialize response
func NewResponse() *Response {
	oneRsp.Do(func() {
		rsp = &Response{}
	})

	rsp.Timestamp = time.Now().Local()
	// clone response
	x := *rsp

	return &x
}
