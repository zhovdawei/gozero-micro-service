package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/logic"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/svc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/types"
)

func userLoginByPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginByPhoneLogic(r.Context(), svcCtx)
		resp, err := l.UserLoginByPhone(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
