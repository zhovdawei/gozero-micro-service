package logic

import (
	"context"
	"github.com/jinzhu/copier"

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
	one, err := l.svcCtx.UserPostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var userPost pb.UserPostResp
	copier.Copy(&userPost, one)
	return &userPost, nil
}
