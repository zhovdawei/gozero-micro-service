package logic

import (
	"context"

	"github.com/zhovdawei/gozero-micro-service/user/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIdQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserIdQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIdQueryLogic {
	return &UserIdQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserIdQueryLogic) UserIdQuery(in *user.IdQueryReq) (*user.UserInfoObj, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserInfoObj{
		Id:       one.Id,
		Name:     one.Name,
		Nickname: "",
		Birthday: "2022-01-01",
		Phone:    one.Phone,
	}, nil
}
