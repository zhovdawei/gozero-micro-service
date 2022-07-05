package logic

import (
	"context"
	"fmt"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/types"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserRpc.UserIdQuery(l.ctx, &user.IdQueryReq{Id: 1})
	if err != nil {
		return nil, err
	}
	fmt.Println("Id = ", userInfo.Id)
	fmt.Println("Name = ", userInfo.Name)
	fmt.Println("Nickname = ", userInfo.Nickname)
	fmt.Println("Phone = ", userInfo.Phone)
	return &types.UserLoginResp{
		Uid: string(userInfo.Id) + "-" + userInfo.Phone,
	}, nil
}
