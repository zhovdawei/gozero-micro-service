package mt

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/logic/mt"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/types"
)

func GetUserPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mt.NewGetUserPostLogic(r.Context(), svcCtx)
		resp, err := l.GetUserPost(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
