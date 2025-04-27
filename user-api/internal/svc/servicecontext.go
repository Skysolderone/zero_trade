package svc

import (
	"trade/user-api/internal/config"
	"trade/user-rpc/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	zclient := zrpc.MustNewClient(c.UserRpc)
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserClient(zclient.Conn()),
	}
}
