package logic

import (
	"context"
	"go-zero-study/rpc/user/users"

	"go-zero-study/api/internal/svc"
	"go-zero-study/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.ReqUserId) (resp *types.CommResp, err error) {
	r, err := l.svcCtx.User.Delete(l.ctx, &users.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: r.Ok, Error: r.Error}, nil
}
