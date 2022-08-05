package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zhovdawei/gozero-micro-service/goods/model"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGoodsListLogic {
	return &QueryGoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGoodsListLogic) QueryGoodsList(in *pb.QueryGoodsListReq) (*pb.GoodsResp, error) {
	// todo: add your logic here and delete this line
	var list []*model.Goods
	var err error
	if in.GoodsName != "" {
		list, err = l.svcCtx.GoodsModel.FindListByGoodsName(l.ctx, in.GoodsName)
	}
	if in.Category != "" {
		list, err = l.svcCtx.GoodsModel.FindListByCategory(l.ctx, in.Category)
	}
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var goodsArr []*pb.GoodsVO
	copier.Copy(&goodsArr, list)
	return &pb.GoodsResp{GoodsList: goodsArr}, nil
}
