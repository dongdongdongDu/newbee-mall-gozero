package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"newbee-mall-gozero/service/admin/api/internal/logic"
	"newbee-mall-gozero/service/admin/api/internal/svc"
)

func GetAdminProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetAdminProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetAdminProfile()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
