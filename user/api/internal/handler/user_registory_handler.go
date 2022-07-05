package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhovdawei/gozero-micro-service/user/api/internal/logic"
	"github.com/zhovdawei/gozero-micro-service/user/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/api/internal/types"
)

func userRegistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRegistoryLogic(r.Context(), svcCtx)
		resp, err := l.UserRegistory(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
