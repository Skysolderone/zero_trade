package logic

import (
	"context"

	"trade/user-rpc/internal/model"
	"trade/user-rpc/internal/svc"
	"trade/user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.InfoReq) (*user.UserInfo, error) {
	userinfo := model.User{}
	err := l.svcCtx.Orm.Model(&model.User{}).Where("id = ?", in.Id).First(&userinfo).Error
	if err != nil {
		return nil, err
	}

	return &user.UserInfo{
		Id:       int32(userinfo.ID),
		Username: userinfo.Username,
		Email:    userinfo.Email,
	}, nil
}
