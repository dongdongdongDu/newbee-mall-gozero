package admin

import (
	"net/http"
	"newbee-mall-gozero/service/admin/api/internal/logic/admin"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"newbee-mall-gozero/service/admin/api/internal/svc"
)

func GetAdminProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAdminProfileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := admin.NewGetAdminProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetAdminProfile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
