//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/xiaohubai/rpc-layout/app/dict/internal/biz"
	"github.com/xiaohubai/rpc-layout/app/dict/internal/conf"
	"github.com/xiaohubai/rpc-layout/app/dict/internal/data"
	"github.com/xiaohubai/rpc-layout/app/dict/internal/server"
	"github.com/xiaohubai/rpc-layout/app/dict/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
