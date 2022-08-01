package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/user/model"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"newbee-mall-gozero/service/user/rpc/internal/svc"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetInfoRequest) (*user.GetInfoResponse, error) {
	// 查询token
	tokenRes, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: in.Token,
	})
	if err != nil {
		logx.Error("token不存在")
		return nil, errors.New("token不存在")
	}
	// 查询用户
	userId := tokenRes.Token.UserId
	userRes, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的用户")
			return nil, errors.New("不存在的用户")
		}
		logx.Error("用户信息获取失败" + err.Error())
		return nil, errors.New("用户信息获取失败" + err.Error())
	}

	return &user.GetInfoResponse{
		NickName:      userRes.NickName,
		LoginName:     userRes.LoginName,
		IntroduceSign: userRes.IntroduceSign,
	}, nil
}
