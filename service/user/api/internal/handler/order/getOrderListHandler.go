package order

import (
	"net/http"

	"newbee-mall-gozero/service/user/api/internal/logic/order"
	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewGetOrderListLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
