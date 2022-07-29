package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"newbee-mall-gozero/common/md5"
	"newbee-mall-gozero/service/user/model"
	"newbee-mall-gozero/service/user/rpc/internal/svc"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateInfoRequest) (*user.UpdateInfoResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(500, "不存在的用户")
		}
		return nil, status.Error(500, "用户信息获取失败"+err.Error())
	}

	res.NickName = in.NickName
	res.IntroduceSign = in.IntroduceSign
	// 输入密码为空，不更改密码
	if in.PasswordMd5 != "" {
		res.PasswordMd5 = md5.MD5V([]byte(in.PasswordMd5))
	}

	// 写入更新
	err = l.svcCtx.UserModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, "用户信息更新"+err.Error())
	}

	return &user.UpdateInfoResponse{
		UserId:        res.UserId,
		NickName:      res.NickName,
		IntroduceSign: res.IntroduceSign,
	}, nil
}
