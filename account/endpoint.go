package account

import "github.com/go-kit/kit/endpoint"

type EndPoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}
