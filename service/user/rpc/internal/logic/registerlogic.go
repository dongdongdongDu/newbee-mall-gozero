package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"newbee-mall-gozero/common/md5"
	"newbee-mall-gozero/service/user/model"
	"time"

	"newbee-mall-gozero/service/user/rpc/internal/svc"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 查询用户名是否存在
	_, err := l.svcCtx.UserModel.FindOneByLoginName(l.ctx, in.LoginName)
	if err == nil {
		return nil, status.Error(500, "存在相同用户名")
	}

	// 创建新用户
	if err == model.ErrNotFound {
		newUser := model.User{
			LoginName:     in.LoginName,
			PasswordMd5:   md5.MD5V([]byte(in.Password)),
			IntroduceSign: "随新所欲，蜂富多彩",
			CreateTime:    time.Now(),
		}

		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, "创建失败:"+err.Error())
		}

		newUser.UserId, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, "创建失败:"+err.Error())
		}

		return &user.RegisterResponse{
			UserId: newUser.UserId,
		}, nil
	}

	return nil, status.Error(500, "创建失败:"+err.Error())
}
