package logic

import (
	"context"

	"trade/user-rpc/internal/model"
	"trade/user-rpc/internal/svc"
	"trade/user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.LoginReq) (*user.Resp, error) {
	logx.Alert(in.Username + in.Password)
	newUser := model.User{
		Username: in.Username,
		Password: in.Password,
	}
	err := l.svcCtx.Orm.Create(&newUser).Error
	if err != nil {
		return &user.Resp{
			Code: 1,
			Msg:  "regist failed",
		}, nil
	}
	return &user.Resp{
		Code: 0,
		Msg:  "regist success",
	}, nil
}
