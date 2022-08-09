package logic

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/user"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginByPhoneLogic {
	return &UserLoginByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginByPhoneLogic) UserLoginByPhone(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// 1.查询用户信息
	rpcResp, err := l.svcCtx.UserRpc.QueryUserByPhone(l.ctx, &user.GetUserByPhone{Phone: req.Phone})
	if err != nil {
		return nil, err
	}
	var user types.UserResp
	copier.Copy(&user, rpcResp)

	// 2.验证密码相等
	data := []byte(req.Password)
	has := md5.Sum(data)
	pwSalt := fmt.Sprintf("%x", has)
	if pwSalt != user.Password {
		return nil, errors.New("密码错误")
	}

	// 3.JWT处理
	return &types.UserLoginResp{Token: "token"}, nil
}
