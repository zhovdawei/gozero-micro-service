package pc

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/logic/pc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/types"
)

func QueryGoodsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsSearchPcReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := pc.NewQueryGoodsListLogic(r.Context(), svcCtx)
		resp, err := l.QueryGoodsList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
