package mt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/user"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPostListLogic {
	return &GetUserPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPostListLogic) GetUserPostList() (resp []types.UserPostResp, err error) {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		l.Logger.Error("json.Number 转化错误，非数字")
		return nil, err
	}
	rpcPostListResp, err := l.svcCtx.UserRpc.QueryUserPostArray(l.ctx, &user.QueryUserPostByUserIdReq{UserId: userId})
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
