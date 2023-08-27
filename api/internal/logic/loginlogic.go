package logic

import (
	"context"
	"go-zero-study/api/internal/svc"
	"go-zero-study/api/internal/types"
	"go-zero-study/rpc/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.ReqUser) (resp *types.RespLogin, err error) {
	// 调用 user rpc 的 login 方法
	r, err := l.svcCtx.User.Login(l.ctx, &users.ReqUser{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &types.RespLogin{Token: r.Token}, nil
}
