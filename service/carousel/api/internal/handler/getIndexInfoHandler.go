package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"newbee-mall-gozero/service/carousel/api/internal/logic"
	"newbee-mall-gozero/service/carousel/api/internal/svc"
)

func GetIndexInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetIndexInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetIndexInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
