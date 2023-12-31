package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-study/api/internal/config"
	"go-zero-study/rpc/user/users"
)

type ServiceContext struct {
	Config config.Config
	// 手动添加
	// users.Users 是 user rpc 服务对外暴露的接口
	User users.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// 手动添加
		//  zrpc.MustNewClient(c.User) 创建了一个 grpc 客户端
		User: users.NewUsers(zrpc.MustNewClient(c.User)),
	}
}
