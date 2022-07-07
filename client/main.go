package main

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	v1 "github.com/xiaohubai/rpc-layout/api/dict/v1"
)

func main() {

	c := consulAPI.DefaultConfig()
	c.Address = "82.156.4.85:8500"
	c.Scheme = "http"
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	endpoint := "discovery://default/rpc-layout"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}

	client := v1.NewDictClient(conn)
	resp, _ := client.Get(context.Background(), &v1.DictRequest{Tag: "ssss"})
	fmt.Println(resp)
}
