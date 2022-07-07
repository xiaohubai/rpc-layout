package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/xiaohubai/rpc-layout/api/dict/v1"
)

// Dict字典序表
type Dict struct {
	Woker  string
	Guide  string
	Actor  string
	Lawyer string
}

type DictRepo interface {
	Get(context.Context) (*Dict, error)
}

type DictUsecase struct {
	repo DictRepo
	log  *log.Helper
}

// NewDictUsecase new a dict usecase.
func NewDictUsecase(repo DictRepo, logger log.Logger) *DictUsecase {
	return &DictUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DictUsecase) GetDict(ctx context.Context) (*pb.DictResponse, error) {
	data, err := uc.repo.Get(ctx)
	if err != nil {
		return nil, err
	}
	result := &pb.DictResponse{
		Woker: data.Woker,
		Guide: "导游",
	}
	return result, err
}
