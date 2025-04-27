package logic

import (
	"context"
	"time"

	"trade/user-rpc/internal/model"
	"trade/user-rpc/internal/svc"
	"trade/user-rpc/pb/user"

	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	logx.Alert(in.Username + in.Password)
	takeUser := model.User{}
	err := l.svcCtx.Orm.Model(model.User{}).Where("username = ? ", in.Username).First(&takeUser).Error
	if err != nil {
		return &user.LoginResp{
			Code: 1,
			Msg:  "login failed",
		}, nil
	}
	return &user.LoginResp{
		Code:  0,
		Msg:   "login success",
		Token: l.generateToekn(takeUser.ID),
	}, nil
}

func (l *LoginLogic) generateToekn(id uint) string {
	claims := jwt.MapClaims{
		"userId": id,
		"exp":    time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
	if err != nil {
		logx.Error(err)
		return ""
	}
	return tokenString
}
