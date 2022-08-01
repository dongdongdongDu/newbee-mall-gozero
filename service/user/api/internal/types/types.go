// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}

type LoginResponse struct {
	UserId     int64  `json:"userId"`
	Token      string `json:"token"`
	UpdateTime int64  `json:"updateTime"`
	ExpireTime int64  `json:"expireTime"`
}

type RegisterRequest struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type RegisterResponse struct {
}

type LogoutRequest struct {
}

type LogoutResponse struct {
}

type UpdateUserInfoRequest struct {
	NickName      string `json:"nickName"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
}

type UpdateUserInfoResponse struct {
	NickName      string `json:"nickName"`
	IntroduceSign string `json:"introduceSign"`
}

type GetUserInfoRequest struct {
}

type GetUserInfoResponse struct {
	NickName      string `json:"nickName"`
	LoginName     string `json:"loginName"`
	IntroduceSign string `json:"introduceSign"`
}

type Response struct {
	ResultCode int         `json:"resultCode"`
	Msg        string      `json:"message"`
	Data       interface{} `json:"data"`
}