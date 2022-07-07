package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/xiaohubai/rpc-layout/app/dict/internal/biz"
)

type dictRepo struct {
	data *Data
	log  *log.Helper
}

// NewDictRepo .
func NewDictRepo(data *Data, logger log.Logger) biz.DictRepo {
	return &dictRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *dictRepo) Get(ctx context.Context) (*biz.Dict, error) {

	return &biz.Dict{}, nil
}
