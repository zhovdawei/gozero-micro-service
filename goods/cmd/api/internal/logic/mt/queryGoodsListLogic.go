package mt

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/pb"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGoodsListLogic {
	return &QueryGoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryGoodsListLogic) QueryGoodsList(req *types.GoodsSearchMtReq) (resp *types.GoodsSearchMtResp, err error) {

	nextSize := req.Size + 1
	rpcResp, err := l.svcCtx.GoodsRpc.QueryGoodsList(l.ctx, &pb.QueryGoodsListReq{
		Category:    req.Category,
		GoodsName:   req.GoodsName,
		LastGoodsId: req.LastGoodsId,
		Size:        int32(nextSize),
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.GoodsSearchMtResp{0, nil}, err
	}
	//
	var goodsList []types.GoodsMtVO
	var isNext = 0
	if rpcResp.GoodsList != nil {
		copier.Copy(&goodsList, rpcResp.GoodsList)
		if len(goodsList) >= nextSize {
			isNext = 1
			goodsList = goodsList[:len(goodsList)-1]
		}
	}
	return &types.GoodsSearchMtResp{isNext, goodsList}, err
}
