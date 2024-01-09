package provider

import (
	"context"

	"github.com/Zainal21/go-bone/pkg/httpclient"
	"github.com/sony/gobreaker"
)

type ExecuteServiceFunc func(options httpclient.RequestOptions) (httpclient.Response, error)
type BackupProvider func(name string, from gobreaker.State, to gobreaker.State)

func SendRequest(reqOption httpclient.RequestOptions) (httpclient.Response, error) {
	response, err := httpclient.Request(reqOption)

	if err != nil {
		return response, err
	}

	return response, nil
}

type CircuitBreaker interface {
	Execute(ctx context.Context, fn ExecuteServiceFunc, options httpclient.RequestOptions) (any, error)
	GetState(ctx context.Context) gobreaker.State
}

type Example interface {
	GetTodos(ctx context.Context) ([]byte, error)
}
