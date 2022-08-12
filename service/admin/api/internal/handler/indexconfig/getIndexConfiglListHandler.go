package indexconfig

import (
	"net/http"

	"newbee-mall-gozero/service/admin/api/internal/logic/indexconfig"
	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetIndexConfiglListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetIndexConfigListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := indexconfig.NewGetIndexConfiglListLogic(r.Context(), svcCtx)
		resp, err := l.GetIndexConfiglList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
