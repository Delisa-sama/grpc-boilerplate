package controllers

import (
	"context"
	"database/sql"

	"github.com/Delisa-sama/logger"
	"github.com/volatiletech/sqlboiler/boil"

	"grpc-boilerplate/api/v1"
	"grpc-boilerplate/models"
)

type MemoController struct {
	DB *sql.DB
}

func (c *MemoController) Add(ctx context.Context, in *v1.MemoAddRequest) (*v1.MemoAddResponse, error) {
	m := &models.Memo{Memo: in.Memo}
	if err := m.Insert(c.DB, boil.Infer()); err != nil {
		logger.Error(err)
		return nil, err
	}

	return &v1.MemoAddResponse{Id: int64(m.ID)}, nil
}

func (c *MemoController) Get(ctx context.Context, in *v1.MemoGetRequest) (*v1.MemoGetResponse, error) {
	m, err := models.Memos(models.MemoWhere.ID.EQ(int(in.Id))).One(c.DB)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &v1.MemoGetResponse{Memo: m.Memo}, nil
}

func (c *MemoController) List(ctx context.Context, in *v1.MemoListRequest) (*v1.MemoListResponse, error) {
	memos, err := models.Memos().All(c.DB)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var m []string
	for _, memo := range memos {
		m = append(m, memo.Memo)
	}

	return &v1.MemoListResponse{Memos: m}, nil
}
