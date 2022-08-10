package mt

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/user"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPostQueryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPostQueryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPostQueryListLogic {
	return &UserPostQueryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPostQueryListLogic) UserPostQueryList(req *types.PostListReq) (resp []types.UserPostResp, err error) {
	rpcPostListResp, err := l.svcCtx.UserRpc.QueryUserPostArray(l.ctx, &user.QueryUserPostByUserIdReq{UserId: req.UserId})
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	if rpcPostListResp == nil {
		return nil, nil
	}
	fmt.Println(*rpcPostListResp)

	var postResult []types.UserPostResp
	copier.Copy(&postResult, rpcPostListResp.UserPosts)

	fmt.Println(postResult)
	return postResult, nil
}
