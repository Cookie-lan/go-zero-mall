package logic

import (
	"context"
	"go-zero-study/rpc/model"

	"go-zero-study/rpc/user/internal/svc"
	"go-zero-study/rpc/user/pb/go-zero-study/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.ReqUser) (*user.CommResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.User{
		Password: in.Password,
		Username: in.Username,
	})
	if err != nil {
		return nil, err
	}

	return &user.CommResp{Ok: true}, nil
}
