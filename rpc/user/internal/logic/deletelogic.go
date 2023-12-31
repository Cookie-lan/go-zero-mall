package logic

import (
	"context"

	"go-zero-study/rpc/user/internal/svc"
	"go-zero-study/rpc/user/pb/go-zero-study/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *user.ReqUserId) (*user.CommResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Model.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.CommResp{Ok: true}, nil
}
