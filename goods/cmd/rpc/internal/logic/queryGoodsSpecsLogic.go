package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGoodsSpecsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGoodsSpecsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGoodsSpecsLogic {
	return &QueryGoodsSpecsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGoodsSpecsLogic) QueryGoodsSpecs(in *pb.QueryGoodsSpecsReq) (*pb.GoodsSpecsResp, error) {
	list, err := l.svcCtx.GoodsSpecModel.FindList(l.ctx, in.GoodsId)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var specs []*pb.GoodsSpecVO
	copier.Copy(&specs, list)
	return &pb.GoodsSpecsResp{GoodsSpecs: specs}, nil
}
