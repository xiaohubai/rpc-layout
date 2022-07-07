package service

import (
	"context"
	"fmt"

	pb "github.com/xiaohubai/rpc-layout/api/dict/v1"
)

func (s *DictService) Get(ctx context.Context, req *pb.DictRequest) (*pb.DictResponse, error) {
	fmt.Println(req.Tag)

	return s.uc.GetDict(ctx)
}
