package logic

import (
	"context"
	"encoding/json"
	"newbee-mall-gozero/common/response"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.Response, err error) {
	// 避免出现未使用的形参提醒 =_=！
	_ = req
	// 从token.Claims中取userId
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetInfoRequest{
		UserId: userId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "未查询到记录",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetUserInfoResponse{
			NickName:      res.NickName,
			LoginName:     res.LoginName,
			IntroduceSign: res.IntroduceSign,
		},
	}, nil
}
