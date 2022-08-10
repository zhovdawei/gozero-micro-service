package login

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/user"
	"time"

	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MtLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMtLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MtLoginLogic {
	return &MtLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MtLoginLogic) MtLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
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

	// 3.获取JWT字符串
	iat := time.Now().Unix()
	authSecret := l.svcCtx.Config.Auth.AccessSecret
	authExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := getJwtToken(authSecret, iat, authExpire, user.UserId)
	if err != nil {
		l.Logger.Error("jwt token 签发失败")
		return nil, err
	}

	return &types.UserLoginResp{
		Token:   token,
		Expire:  iat + authExpire*2,
		Refresh: iat + authExpire,
	}, nil
}

// 生成jwt token
func getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	claims["userId"] = userId
	//token := jwt.New(jwt.SigningMethodHS256)
	//token.Claims = claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
