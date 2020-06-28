package controllers

import (
	"context"
	"database/sql"
	"grpc-boilerplate/api"
)

type EchoController struct {
	DB *sql.DB
}

func (c *EchoController) Echo(ctx context.Context, in *api.EchoRequest) (*api.EchoResponse, error) {
	return &api.EchoResponse{Message: in.Message}, nil
}
