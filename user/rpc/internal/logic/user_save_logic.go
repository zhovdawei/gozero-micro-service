package logic

import (
	"context"

	"github.com/zhovdawei/gozero-micro-service/user/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSaveLogic {
	return &UserSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSaveLogic) UserSave(in *user.UserInfoObj) (*user.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &user.CommonResp{}, nil
}
