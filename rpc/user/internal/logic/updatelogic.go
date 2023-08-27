package logic

import (
	"context"
	"go-zero-study/rpc/model"

	"go-zero-study/rpc/user/internal/svc"
	"go-zero-study/rpc/user/pb/go-zero-study/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *user.User) (*user.CommResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Model.Update(l.ctx, &model.User{
		Password: in.Password,
		Id:       in.Id,
		Username: in.Username,
	})
	if err != nil {
		return nil, err
	}
	return &user.CommResp{Ok: true}, nil
}
