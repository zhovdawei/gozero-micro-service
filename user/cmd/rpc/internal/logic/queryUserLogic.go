package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserLogic {
	return &QueryUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserLogic) QueryUser(in *pb.QueryUserByIdReq) (*pb.UserResp, error) {
	// todo: add your logic here and delete this line
	one,err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var user pb.UserResp
	copier.Copy(&user,one)
	return &user, nil
}
