package handler

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"newbee-mall-gozero/service/user/api/internal/logic"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		// 把token传入Context中
		logx.Info("r.Header['token']:", r.Header.Get("token"))
		token := r.Header.Get("token")
		ctx := context.WithValue(r.Context(), "token", token)

		l := logic.NewGetUserInfoLogic(ctx, svcCtx)
		resp, err := l.GetUserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
