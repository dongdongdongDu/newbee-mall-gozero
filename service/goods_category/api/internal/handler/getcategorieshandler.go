package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"newbee-mall-gozero/service/goods_category/api/internal/logic"
	"newbee-mall-gozero/service/goods_category/api/internal/svc"
)

func GetCategoriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetCategoriesLogic(r.Context(), svcCtx)
		resp, err := l.GetCategories()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
