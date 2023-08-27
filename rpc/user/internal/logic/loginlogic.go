package logic

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"go-zero-study/rpc/user/internal/svc"
	"go-zero-study/rpc/user/pb/go-zero-study/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const secretKey = "go-zero-demo"

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.ReqUser) (*user.RespLogin, error) {
	one, err := l.svcCtx.Model.FindByName(in.Username)
	if err != nil {
		return nil, errors.Wrapf(err, "find user by name: %s", in.Username)
	}

	if one.Password != in.Password {
		return nil, fmt.Errorf("user or password is invalid")
	}

	token := GenTokenByHmac(one.Username, secretKey)
	return &user.RespLogin{Token: token}, nil
}

func HmacCrypto(s, key string) string {
	hc := hmac.New(sha1.New, []byte(key))
	token := hc.Sum([]byte(s))
	return hex.EncodeToString(token)
}

func GenTokenByHmac(s, key string) string {
	return fmt.Sprintf("%s,%s", s, HmacCrypto(s, key))
}
