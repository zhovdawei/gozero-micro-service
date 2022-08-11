package mt

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/goods"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsAttrsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsAttrsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsAttrsLogic {
	return &GetGoodsAttrsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsAttrsLogic) GetGoodsAttrs(req *types.GoodsAttrsMtReq) (resp *types.GoodsAttrsMtResp, err error) {
	rpcResp, err := l.svcCtx.GoodsRpc.QueryGoodsAttrs(l.ctx, &goods.QueryGoodsAttrsReq{GoodsId: req.GoodsId})
	if err != nil {
		l.Logger.Errorf("%+v", errors.New(err.Error()))
		return nil, err
	}
	if rpcResp == nil || rpcResp.GoodsAttrs == nil {
		return nil, nil
	}
	var attrs []*types.GoodsAttrVO
	copier.Copy(&attrs, rpcResp.GoodsAttrs)
	for _, attr := range attrs {
		attrVal := attr.AttrVals
		var vals []string
		err := json.Unmarshal([]byte(attrVal), &vals)
		if err != nil {
			l.Logger.Errorf("%+v", errors.New("json反序列化失败，jsonstring="+attrVal))
			return nil, err
		}
		attr.AttrValList = vals
	}
	return &types.GoodsAttrsMtResp{GoodsAttrs: attrs}, nil
}
