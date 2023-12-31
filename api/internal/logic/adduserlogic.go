package logic

import (
	"context"
	"go-zero-study/rpc/user/users"

	"go-zero-study/api/internal/svc"
	"go-zero-study/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.ReqUser) (resp *types.CommResp, err error) {
	r, err := l.svcCtx.User.Create(l.ctx, &users.ReqUser{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &types.CommResp{Ok: r.Ok, Error: r.Error}, nil
}
