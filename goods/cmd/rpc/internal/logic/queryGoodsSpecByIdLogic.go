package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGoodsSpecByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGoodsSpecByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGoodsSpecByIdLogic {
	return &QueryGoodsSpecByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGoodsSpecByIdLogic) QueryGoodsSpecById(in *pb.QueryGoodsSpecByIdReq) (*pb.GoodsSpecVO, error) {
	one, err := l.svcCtx.GoodsSpecModel.FindOne(l.ctx, in.GoodsSpecId)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var goodsSpec pb.GoodsSpecVO
	copier.Copy(&goodsSpec, one)
	return &goodsSpec, nil
}
