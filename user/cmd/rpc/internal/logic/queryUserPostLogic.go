package logic

import (
	"context"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserPostLogic {
	return &QueryUserPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserPostLogic) QueryUserPost(in *pb.QueryUserPostReq) (*pb.UserPostResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserPostResp{}, nil
}
