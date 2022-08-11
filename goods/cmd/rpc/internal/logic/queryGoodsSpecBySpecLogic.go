package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGoodsSpecBySpecLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGoodsSpecBySpecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGoodsSpecBySpecLogic {
	return &QueryGoodsSpecBySpecLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGoodsSpecBySpecLogic) QueryGoodsSpecBySpec(in *pb.QueryGoodsSpecBySpecReq) (*pb.GoodsSpecVO, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.GoodsSpecModel.QueryGoodsSpecBySpec(l.ctx, in.GoodsId, in.Specs)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var goodsSpec pb.GoodsSpecVO
	copier.Copy(&goodsSpec, one)
	return &goodsSpec, nil
}
