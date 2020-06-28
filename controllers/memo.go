package controllers

import (
	"context"
	"database/sql"
	"github.com/Delisa-sama/logger"
	"github.com/volatiletech/sqlboiler/boil"
	"grpc-boilerplate/api"
	"grpc-boilerplate/models"
)

type MemoController struct {
	DB *sql.DB
}

func (c *MemoController) Add(ctx context.Context, in *api.MemoAddRequest) (*api.MemoAddResponse, error) {
	m := &models.Memo{Memo: in.Memo}
	if err := m.Insert(c.DB, boil.Infer()); err != nil {
		logger.Error(err)
		return nil, err
	}

	return &api.MemoAddResponse{Id: int64(m.ID)}, nil
}

func (c *MemoController) Get(ctx context.Context, in *api.MemoGetRequest) (*api.MemoGetResponse, error) {
	m, err := models.Memos(models.MemoWhere.ID.EQ(int(in.Id))).One(c.DB)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &api.MemoGetResponse{Memo: m.Memo}, nil
}

func (c *MemoController) List(ctx context.Context, in *api.MemoListRequest) (*api.MemoListResponse, error) {
	memos, err := models.Memos().All(c.DB)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var m []string
	for _, memo := range memos {
		m = append(m, memo.Memo)
	}

	return &api.MemoListResponse{Memos: m}, nil
}
