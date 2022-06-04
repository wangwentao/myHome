package remote

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Reply string `json:"reply"`
}

func MakeHelloEndPoint(s IHelloService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req, ok := request.(HelloRequest)
		if !ok {
			return HelloResponse{}, nil
		}

		resp := s.Hello(req.Name)
		return HelloResponse{Reply: resp}, nil

	}
}
