package controllers

import (
	"context"

	"grpc-boilerplate/api/v1"
)

type EchoController struct{}

func (c *EchoController) Echo(ctx context.Context, in *v1.EchoRequest) (*v1.EchoResponse, error) {
	return &v1.EchoResponse{Message: in.Message}, nil
}
