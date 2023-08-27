package logic

import (
	"context"
	"go-zero-study/rpc/user/users"

	"go-zero-study/api/internal/svc"
	"go-zero-study/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.ReqUserId) (resp *types.User, err error) {
	r, err := l.svcCtx.User.Get(l.ctx, &users.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}

	resp = &types.User{
		Id:       r.Id,
		Username: r.Username,
		Password: r.Password,
	}
	return resp, nil
}
