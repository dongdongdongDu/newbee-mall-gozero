package user

import (
	"context"
	"encoding/json"
	"newbee-mall-gozero/service/user/rpc/user"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoRequest) (resp *types.UpdateUserInfoResponse, err error) {
	// 从token.Claims中取userId
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	// 更新信息
	res, err := l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &user.UpdateInfoRequest{
		UserId:        userId,
		NickName:      req.NickName,
		PasswordMd5:   req.PasswordMd5,
		IntroduceSign: req.IntroduceSign,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateUserInfoResponse{
		NickName:      res.NickName,
		IntroduceSign: res.IntroduceSign,
	}, nil
}
