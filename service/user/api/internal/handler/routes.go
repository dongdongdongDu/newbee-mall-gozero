// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	goodsInfo "newbee-mall-gozero/service/user/api/internal/handler/goodsInfo"
	user "newbee-mall-gozero/service/user/api/internal/handler/user"
	userAddress "newbee-mall-gozero/service/user/api/internal/handler/userAddress"
	"newbee-mall-gozero/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/info",
					Handler: user.UpdateUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/goods/detail/:id",
					Handler: goodsInfo.GetGoodsDetailHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/address",
					Handler: userAddress.SaveAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/address",
					Handler: userAddress.UpdateAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/address/:addressId",
					Handler: userAddress.DeleteAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/address",
					Handler: userAddress.GetMyAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/address/:addressId",
					Handler: userAddress.GetAddressByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/address/default",
					Handler: userAddress.GetDefaultAddressHandler(serverCtx),
				},
			}...,
		),
	)
}
