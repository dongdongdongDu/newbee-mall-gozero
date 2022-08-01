package logic

import (
	"context"
	"errors"
	"time"

	"newbee-mall-gozero/common/jwt"
	"newbee-mall-gozero/service/admin_token/model"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/admin_token/rpc/internal/svc"

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

func (l *GenerateTokenLogic) GenerateToken(in *admintoken.GenerateTokenRequest) (*admintoken.GenerateTokenResponse, error) {
	// 生成token
	token := jwt.GetNewToken(time.Now().UnixNano()/1e6, int(in.AdminUserId))
	nowDate := time.Now()
	// 48小时过期
	expireTime, _ := time.ParseDuration("48h")
	expireDate := nowDate.Add(expireTime)

	// 查询用户是否已存在token
	res, err := l.svcCtx.AdminTokenModel.FindOne(l.ctx, in.AdminUserId)
	if err == nil {
		// 有，更新
		logx.Error("生成新的adminTokenModel")
		res.Token = token
		res.UpdateTime = nowDate
		res.ExpireTime = expireDate

		// 写入更新
		err = l.svcCtx.AdminTokenModel.Update(l.ctx, res)
		if err != nil {
			logx.Error("adminToken更新失败" + err.Error())
			return nil, errors.New("adminToken更新失败" + err.Error())
		}

		return &admintoken.GenerateTokenResponse{
			AdminToken: &admintoken.AdminToken{
				AdminUserId: res.AdminUserId,
				Token:       res.Token,
				UpdateTime:  res.UpdateTime.Unix(),
				ExpireTime:  res.ExpireTime.Unix(),
			},
		}, nil
	} else {
		// 没有
		logx.Error("生成新的adminTokenModel")
		// 创建新tokenModel
		if err == model.ErrNotFound {
			newToken := model.TbNewbeeMallAdminUserToken{
				AdminUserId: in.AdminUserId,
				Token:       token,
				UpdateTime:  nowDate,
				ExpireTime:  expireDate,
			}

			_, err := l.svcCtx.AdminTokenModel.Insert(l.ctx, &newToken)
			if err != nil {
				logx.Error("创建token失败：" + err.Error())
				return nil, err
			}

			return &admintoken.GenerateTokenResponse{
				AdminToken: &admintoken.AdminToken{
					AdminUserId: newToken.AdminUserId,
					Token:       newToken.Token,
					UpdateTime:  newToken.UpdateTime.Unix(),
					ExpireTime:  newToken.ExpireTime.Unix(),
				},
			}, nil
		}
		logx.Error("创建token失败：" + err.Error())

		return nil, err
	}
}
