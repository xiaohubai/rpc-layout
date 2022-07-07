package service

import (
	"github.com/google/wire"
	pb "github.com/xiaohubai/rpc-layout/api/dict/v1"
	"github.com/xiaohubai/rpc-layout/app/dict/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDictService)

type DictService struct {
	pb.UnimplementedDictServer
	uc *biz.DictUsecase
}

func NewDictService(uc *biz.DictUsecase) *DictService {
	return &DictService{uc: uc}
}
