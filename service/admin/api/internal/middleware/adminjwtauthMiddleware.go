package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"time"
)

type AdminJwtAuthMiddleware struct {
	adminToken admintoken.Admintoken
}

func NewAdminJwtAuthMiddleware(t admintoken.Admintoken) *AdminJwtAuthMiddleware {
	return &AdminJwtAuthMiddleware{
		adminToken: t,
	}
}

func (m *AdminJwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求中获取token
		token := r.Header.Get("token")
		// 没有token
		if token == "" {
			httpx.OkJson(w, map[string]interface{}{
				"resultCode": response.ERROR,
				"msg":        "未登录或非法访问",
				"data":       nil,
			})
			logx.Info("未登录或非法访问")
			return
		}
		// 有token
		userToken, err := m.adminToken.GetExistToken(r.Context(), &admintoken.GetExistTokenRequest{
			Token: token,
		})
		// 不匹配
		if err != nil {
			httpx.OkJson(w, map[string]interface{}{
				"resultCode": response.ERROR,
				"msg":        "未登录或非法访问",
				"data":       nil,
			})
			logx.Info("未登录或非法访问")
			return
		}

		// 判断是否有效
		if time.Now().Unix() > userToken.AdminToken.ExpireTime {
			// 无效
			httpx.OkJson(w, map[string]interface{}{
				"resultCode": response.ERROR,
				"msg":        "授权已过期",
				"data":       nil,
			})
			// 删除这个token
			_, err := m.adminToken.DeleteToken(r.Context(), &admintoken.DeleteTokenRequest{
				Token: token,
			})
			if err != nil {
				logx.Error("删除token出错")
				return
			}
			return
		}
		// 有效的token，通过
		next(w, r)
	}
}
