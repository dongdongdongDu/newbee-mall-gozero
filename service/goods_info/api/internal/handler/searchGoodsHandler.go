package handler

import (
	"net/http"

	"newbee-mall-gozero/service/goods_info/api/internal/logic"
	"newbee-mall-gozero/service/goods_info/api/internal/svc"
	"newbee-mall-gozero/service/goods_info/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchGoodsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchGoodsLogic(r.Context(), svcCtx)
		resp, err := l.SearchGoods(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
