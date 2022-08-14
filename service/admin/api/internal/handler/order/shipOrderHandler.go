package order

import (
	"net/http"

	"newbee-mall-gozero/service/admin/api/internal/logic/order"
	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShipOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewShipOrderLogic(r.Context(), svcCtx)
		resp, err := l.ShipOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
