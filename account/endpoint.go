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
			req := request.(CreateUserRequest)
			ok, err := s.CreateUser(ctx, req.Email, req.Password)
			return CreateUserResponse{Ok: ok}, err
		},
		GetUser: func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(GetUserRequest)
			email, err := s.GetUser(ctx, req.ID)
			return GetUserResponse{
				Email: email,
			}, err
		},
	}
}
