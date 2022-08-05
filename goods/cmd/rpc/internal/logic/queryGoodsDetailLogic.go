package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGoodsDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGoodsDetailLogic {
	return &QueryGoodsDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGoodsDetailLogic) QueryGoodsDetail(in *pb.QueryGoodsDetailReq) (*pb.GoodsDetailResp, error) {
	// 1.先查goods
	one, err := l.svcCtx.GoodsModel.FindOne(l.ctx, in.GoodsId)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var goods pb.GoodsVO
	copier.Copy(&goods, one)

	// 2.再查goodsAttr
	list, err := l.svcCtx.GoodsAttrsModel.FindList(l.ctx, in.GoodsId)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var details []*pb.GoodsAttrsVO
	copier.Copy(&details, list)

	return &pb.GoodsDetailResp{
		Goods:      &goods,
		GoodsAttrs: details,
	}, nil
}
