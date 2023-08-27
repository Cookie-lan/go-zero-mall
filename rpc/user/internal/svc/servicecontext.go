package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-study/rpc/model"
	"go-zero-study/rpc/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
