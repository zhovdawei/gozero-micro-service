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

type GetUserPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPostLogic {
	return &GetUserPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPostLogic) GetUserPost(req *types.PostReq) (resp *types.UserPostResp, err error) {
	rpcPostResp, err := l.svcCtx.UserRpc.QueryUserPost(l.ctx, &user.QueryUserPostReq{PostId: req.PostId})
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	if rpcPostResp == nil {
		return nil, nil
	}
	fmt.Println(*rpcPostResp)

	var postResult types.UserPostResp
	copier.Copy(&postResult, rpcPostResp)

	fmt.Println(postResult)
	return &postResult, nil
}
