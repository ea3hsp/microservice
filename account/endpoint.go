package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) EndPoints {
	return EndPoints{
		CreateUser: func(ctx context.Context, request interface{}) (interface{}, error) {
		},
		GetUser: func(ctx context.Context, request interface{}) (interface{}, error) {
		},
	}
}
