package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGoodsAttrsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGoodsAttrsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGoodsAttrsLogic {
	return &QueryGoodsAttrsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGoodsAttrsLogic) QueryGoodsAttrs(in *pb.QueryGoodsAttrsReq) (*pb.GoodsAttrsResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.GoodsAttrsModel.FindList(l.ctx, in.GoodsId)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	var details []*pb.GoodsAttrsVO
	copier.Copy(&details, list)
	return &pb.GoodsAttrsResp{GoodsAttrs: details}, nil
}
