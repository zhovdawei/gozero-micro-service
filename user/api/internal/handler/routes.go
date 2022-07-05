// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/zhovdawei/gozero-micro-service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/registory",
				Handler: userRegistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: userLoginHandler(serverCtx),
			},
		},
	)
}
