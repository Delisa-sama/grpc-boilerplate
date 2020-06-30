package controllers

import (
	"context"

	"grpc-boilerplate/api"
)

type EchoController struct{}

func (c *EchoController) Echo(ctx context.Context, in *api.EchoRequest) (*api.EchoResponse, error) {
	return &api.EchoResponse{Message: in.Message}, nil
}
