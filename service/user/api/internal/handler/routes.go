// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"newbee-mall-gozero/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/info",
				Handler: UpdateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/logout",
				Handler: LogoutHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
