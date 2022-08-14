// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	admin "newbee-mall-gozero/service/admin/api/internal/handler/admin"
	carousel "newbee-mall-gozero/service/admin/api/internal/handler/carousel"
	category "newbee-mall-gozero/service/admin/api/internal/handler/category"
	goodsInfo "newbee-mall-gozero/service/admin/api/internal/handler/goodsInfo"
	indexconfig "newbee-mall-gozero/service/admin/api/internal/handler/indexconfig"
	order "newbee-mall-gozero/service/admin/api/internal/handler/order"
	user "newbee-mall-gozero/service/admin/api/internal/handler/user"
	"newbee-mall-gozero/service/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/admin/login",
				Handler: admin.AdminLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/admin/profile",
					Handler: admin.GetAdminProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/admin/name",
					Handler: admin.UpdateAdminNameHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/admin/password",
					Handler: admin.UpdateAdminPwdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/admin/logout",
					Handler: admin.AdminLogoutHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/admin/users",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/admin/users/:lockStatus",
					Handler: user.LockUserHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/goods",
					Handler: goodsInfo.AddGoodsInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/goods",
					Handler: goodsInfo.UpdateGoodsInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/goods/status/:status",
					Handler: goodsInfo.AlterGoodsStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/goods/delete",
					Handler: goodsInfo.DeleteGoodsInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/goods/:id",
					Handler: goodsInfo.GetGoodsInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/goods/list",
					Handler: goodsInfo.GetGoodsListHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/categories",
					Handler: category.AddCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/categories",
					Handler: category.UpdateCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/categories",
					Handler: category.DeleteCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/categories",
					Handler: category.GetCategoryListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/categories/:id",
					Handler: category.GetCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/categories4Select",
					Handler: category.GetCategoryForSelectHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/carousels",
					Handler: carousel.AddCarouselHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/carousels",
					Handler: carousel.UpdateCarouselHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/carousels",
					Handler: carousel.DeleteCarouselHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/carousels/:id",
					Handler: carousel.GetCarouselHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/carousels",
					Handler: carousel.GetCarouselListHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/indexConfigs",
					Handler: indexconfig.AddIndexConfigHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/indexConfigs",
					Handler: indexconfig.UpdateIndexConfigHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/indexConfigs",
					Handler: indexconfig.DeleteIndexConfigHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/indexConfigs/:id",
					Handler: indexconfig.GetIndexConfigHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/indexConfigs",
					Handler: indexconfig.GetIndexConfiglListHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminJwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/orders/checkDone",
					Handler: order.ShipOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/orders/checkOut",
					Handler: order.CheckOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/orders/close",
					Handler: order.CloseOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/orders/:orderId",
					Handler: order.GetOrderByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/orders",
					Handler: order.GetOrdersListHandler(serverCtx),
				},
			}...,
		),
	)
}
