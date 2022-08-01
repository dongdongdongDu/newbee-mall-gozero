package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"
	"time"
)

type UserJwtAuthMiddleware struct {
	userToken usertoken.Usertoken
}

func NewUserJwtAuthMiddleware(t usertoken.Usertoken) *UserJwtAuthMiddleware {
	return &UserJwtAuthMiddleware{
		userToken: t,
	}
}

func (m *UserJwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求中获取token
		token := r.Header.Get("token")
		// 没有token
		if token == "" {
			httpx.OkJson(w, map[string]interface{}{
				"resultCode": response.UNLOGIN,
				"msg":        "未登录！",
				"data":       nil,
			})
			logx.Info("未登录")
			return
		}
		// 有token
		userToken, err := m.userToken.GetExistToken(r.Context(), &usertoken.GetExistTokenRequest{
			Token: token,
		})
		// 不匹配
		if err != nil {
			httpx.OkJson(w, map[string]interface{}{
				"resultCode": response.UNLOGIN,
				"msg":        "未登录！",
				"data":       nil,
			})
			logx.Info("未登录")
			return
		}
		logx.Info("现在：", time.Now().Unix())
		logx.Info("过期：", userToken.Token.ExpireTime)
		// 判断是否有效
		if time.Now().Unix() > userToken.Token.ExpireTime {
			// 无效
			httpx.OkJson(w, map[string]interface{}{
				"resultCode": response.ERROR,
				"msg":        "授权已过期",
				"data":       nil,
			})
			// 删除这个token
			_, err := m.userToken.DeleteToken(r.Context(), &usertoken.DeleteTokenRequest{
				Token: token,
			})
			if err != nil {
				logx.Error("删除token出错")
				return
			}
			return
		}
		// 有效的token，通过
		claims := make(map[string]interface{})
		claims["token"] = token
		next(w, r)
	}
}
