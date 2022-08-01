package tests

import (
	"fmt"
	"newbee-mall-gozero/service/user_token/model"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"
	"testing"
	"time"
)

func TestCopier(t *testing.T) {
	var userId int64 = 1
	token := "abcclasdkjkahfkhagkgk"
	nowDate := time.Now()
	// 48小时过期
	expireTime, _ := time.ParseDuration("48h")
	expireDate := nowDate.Add(expireTime)

	newToken := model.TbNewbeeMallUserToken{
		UserId:     userId,
		Token:      token,
		UpdateTime: nowDate,
		ExpireTime: expireDate,
	}
	fmt.Println(newToken)

	Token := usertoken.Token{
		UserId:     newToken.UserId,
		Token:      newToken.Token,
		UpdateTime: newToken.UpdateTime.Unix(),
		ExpireTime: newToken.ExpireTime.Unix(),
	}
	fmt.Println(Token.UserId)
	fmt.Println(Token.Token)
	fmt.Println(Token.UpdateTime)
	fmt.Println(Token.ExpireTime)

	//var tokenRes *usertoken.Token
	//err := copier.Copy(&tokenRes, newToken)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(tokenRes.UserId)
	//fmt.Println(tokenRes.Token)
	//fmt.Println(tokenRes.UpdateTime)
	//fmt.Println(tokenRes.ExpireTime)

}
