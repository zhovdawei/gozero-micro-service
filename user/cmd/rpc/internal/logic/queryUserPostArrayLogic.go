package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserPostArrayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserPostArrayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserPostArrayLogic {
	return &QueryUserPostArrayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryUserPostArrayLogic) QueryUserPostArray(in *pb.QueryUserPostByUserIdReq) (*pb.UserPostArrayResp, error) {
	list, err := l.svcCtx.UserPostModel.FindList(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	var posts []*pb.UserPostResp
	copier.Copy(&posts, list)
	return &pb.UserPostArrayResp{UserPosts: posts}, nil
}
