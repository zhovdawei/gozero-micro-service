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

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserResp, err error) {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		l.Logger.Error("json.Number 转化错误，非数字")
	}
	rpcUserResp, err := l.svcCtx.UserRpc.QueryUser(l.ctx, &user.QueryUserByIdReq{UserId: userId})
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	if rpcUserResp == nil {
		return nil, nil
	}
	fmt.Println(*rpcUserResp)
	var userResult types.UserResp

	copier.Copy(&userResult, rpcUserResp)
	fmt.Println(userResult)
	return &userResult, nil
}
