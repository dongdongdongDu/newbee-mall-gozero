package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/common/jwt"
	"time"

	"newbee-mall-gozero/service/user_token/model"
	"newbee-mall-gozero/service/user_token/rpc/internal/svc"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *usertoken.GenerateTokenRequest) (*usertoken.GenerateTokenResponse, error) {
	// 生成token
	token := jwt.GetNewToken(time.Now().UnixNano()/1e6, int(in.UserId))
	nowDate := time.Now()
	// 48小时过期
	expireTime, _ := time.ParseDuration("48h")
	expireDate := nowDate.Add(expireTime)

	// 查询用户是否已存在token
	res, err := l.svcCtx.TokenModel.FindOne(l.ctx, in.UserId)
	if err == nil {
		// 有，更新
		logx.Error("生成新的tokenModel")
		res.Token = token
		res.UpdateTime = nowDate
		res.ExpireTime = expireDate

		// 写入更新
		err = l.svcCtx.TokenModel.Update(l.ctx, res)
		if err != nil {
			logx.Error("用户token更新失败" + err.Error())
			return nil, errors.New("用户token更新失败" + err.Error())
		}

		return &usertoken.GenerateTokenResponse{
			Token: &usertoken.Token{
				UserId:     res.UserId,
				Token:      res.Token,
				UpdateTime: res.UpdateTime.Unix(),
				ExpireTime: res.ExpireTime.Unix(),
			},
		}, nil
	} else {
		// 没有
		logx.Error("生成新的tokenModel")
		// 创建新tokenModel
		if err == model.ErrNotFound {
			newToken := model.TbNewbeeMallUserToken{
				UserId:     in.UserId,
				Token:      token,
				UpdateTime: nowDate,
				ExpireTime: expireDate,
			}

			_, err := l.svcCtx.TokenModel.Insert(l.ctx, &newToken)
			if err != nil {
				logx.Error("创建token失败：" + err.Error())
				return nil, err
			}

			return &usertoken.GenerateTokenResponse{
				Token: &usertoken.Token{
					UserId:     newToken.UserId,
					Token:      newToken.Token,
					UpdateTime: newToken.UpdateTime.Unix(),
					ExpireTime: newToken.ExpireTime.Unix(),
				},
			}, nil
		}
		logx.Error("创建token失败：" + err.Error())

		return nil, err
	}
}
