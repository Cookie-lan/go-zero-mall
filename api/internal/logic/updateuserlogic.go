package logic

import (
	"context"
	"go-zero-study/rpc/user/users"

	"go-zero-study/api/internal/svc"
	"go-zero-study/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.ReqUpdateUser) (resp *types.CommResp, err error) {
	r, err := l.svcCtx.User.Update(l.ctx, &users.User{Id: req.Id, Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: r.Ok, Error: r.Error}, nil
}
