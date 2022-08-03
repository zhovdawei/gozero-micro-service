package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/user"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserQueryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserQueryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryByIdLogic {
	return &UserQueryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserQueryByIdLogic) UserQueryById(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	rpcUserResp, err := l.svcCtx.UserRpc.QueryUser(l.ctx,&user.QueryUserByIdReq{UserId: req.UserId})
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	if rpcUserResp == nil {
		return nil,nil
	}
	fmt.Println(*rpcUserResp)
	var userResult types.UserResp

	copier.Copy(&userResult,rpcUserResp)
	fmt.Println(userResult)
	return &userResult,nil
}
