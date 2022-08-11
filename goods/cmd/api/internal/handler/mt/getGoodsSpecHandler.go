package mt

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/logic/mt"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/types"
)

func GetGoodsSpecHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsSpecMtReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mt.NewGetGoodsSpecLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodsSpec(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
