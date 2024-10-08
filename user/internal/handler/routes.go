// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"iot-platform/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/add",
				Handler: UseraddHandler(serverCtx),
			},
		},
	)
}
