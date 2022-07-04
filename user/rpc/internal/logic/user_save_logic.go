package logic

import (
	"context"

	"github.com/zhovdawei/gozero-micro-service/user/model"
	"github.com/zhovdawei/gozero-micro-service/user/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSaveLogic {
	return &UserSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSaveLogic) UserSave(in *user.UserInfoObj) (*user.CommonResp, error) {
	// todo: add your logic here and delete this line

	if in != nil {

		if in.Id == 0 {
			_, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
				Name:     in.Name,
				Phone:    in.Phone,
				Password: "123456",
				Gender:   1,
			})

			if err != nil {
				return nil, err
			}
		} else {
			err := l.svcCtx.UserModel.Update(l.ctx, &model.User{
				Id:       in.Id,
				Name:     in.Name,
				Phone:    in.Phone,
				Password: "123456",
				Gender:   1,
			})

			if err != nil {
				return nil, err
			}
		}
	}

	return &user.CommonResp{Result: true}, nil
}
