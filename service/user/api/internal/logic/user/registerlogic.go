package user

import (
	"context"
	"google.golang.org/grpc/status"
	"newbee-mall-gozero/common/verify"
	"newbee-mall-gozero/service/user/rpc/user"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// 校验输入格式
	if err := verify.Verify(*req, verify.MallUserRegisterVerify); err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 注册
	_, err = l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		LoginName: req.LoginName,
		Password:  req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{}, nil
}
