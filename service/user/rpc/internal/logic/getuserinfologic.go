package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"newbee-mall-gozero/service/user/model"

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
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(500, "不存在的用户")
		}
		return nil, status.Error(500, "用户信息获取失败"+err.Error())
	}

	return &user.GetInfoResponse{
		UserId:        res.UserId,
		NickName:      res.NickName,
		LoginName:     res.LoginName,
		IntroduceSign: res.IntroduceSign,
	}, nil
}
