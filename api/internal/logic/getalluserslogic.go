package logic

import (
	"context"
	"go-zero-study/rpc/user/users"

	"go-zero-study/api/internal/svc"
	"go-zero-study/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUsersLogic {
	return &GetAllUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUsersLogic) GetAllUsers() (resp []types.User, err error) {
	r, err := l.svcCtx.User.GetAll(l.ctx, &users.ReqGetAll{})
	if err != nil {
		return nil, err
	}

	for _, v := range r.Users {
		resp = append(resp, types.User{Id: v.Id, Username: v.Username, Password: v.Password})
	}

	return resp, nil
}
