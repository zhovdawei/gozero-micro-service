package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserByPhoneLogic {
	return &QueryUserByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserByPhoneLogic) QueryUserByPhone(in *pb.GetUserByPhone) (*pb.UserResp, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.UserModel.FindUserByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	var user pb.UserResp
	copier.Copy(&user, one)
	return &user, nil
}
