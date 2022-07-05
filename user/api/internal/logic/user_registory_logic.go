package logic

import (
	"context"

	"github.com/zhovdawei/gozero-micro-service/user/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegistoryLogic {
	return &UserRegistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegistoryLogic) UserRegistory(req *types.UserRegistoryReq) (resp *types.UserRegistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
