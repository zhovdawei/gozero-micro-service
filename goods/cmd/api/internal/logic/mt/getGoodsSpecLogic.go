package mt

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/goods"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsSpecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsSpecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsSpecLogic {
	return &GetGoodsSpecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsSpecLogic) GetGoodsSpec(req *types.GoodsSpecMtReq) (resp *types.GoodsSpecMtResp, err error) {
	// todo: add your logic here and delete this line
	var rpcResp *goods.GoodsSpecVO
	if req.GoodsSpecId != 0 {
		rpcResp, err = l.svcCtx.GoodsRpc.QueryGoodsSpecById(l.ctx, &goods.QueryGoodsSpecByIdReq{GoodsSpecId: req.GoodsSpecId})
	} else {
		rpcResp, err = l.svcCtx.GoodsRpc.QueryGoodsSpecBySpec(l.ctx, &goods.QueryGoodsSpecBySpecReq{GoodsId: req.GoodsId, Specs: req.Spesc})
	}
	if err != nil {
		l.Logger.Errorf("%+v", errors.New(err.Error()))
		return nil, err
	}
	fmt.Println("GoodsSpec ===>>> ", rpcResp)
	var spec types.GoodsSpecMtResp
	copier.Copy(&spec, rpcResp)
	return &spec, nil
}
