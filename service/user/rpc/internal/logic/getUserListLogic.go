package logic

import (
	"context"
	"errors"
	"math"
	"newbee-mall-gozero/service/user/rpc/internal/svc"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *user.GetUserListRequest) (*user.GetUserListResponse, error) {
	// 分页
	limit := in.PageSize
	offset := in.PageSize * (in.PageNumber - 1)
	// 查询用户
	usersRes, err := l.svcCtx.UserModel.TakePart(l.ctx, limit, offset)
	if err != nil {
		logx.Error("用户列表获取失败" + err.Error())
		return nil, errors.New("用户列表获取失败" + err.Error())
	}
	// 查询总数
	totalCount, err := l.svcCtx.UserModel.CountAll(l.ctx)
	if err != nil {
		logx.Error("用户总数获取失败" + err.Error())
		return nil, errors.New("用户总数获取失败" + err.Error())
	}

	usersList := make([]*user.UserModel, 0)
	for _, u := range usersRes {
		usersList = append(usersList, &user.UserModel{
			UserId:        u.UserId,
			NickName:      u.NickName,
			LoginName:     u.LoginName,
			PasswordMd5:   u.PasswordMd5,
			IntroduceSign: u.IntroduceSign,
			IsDeleted:     u.IsDeleted,
			LockedFlag:    u.LockedFlag,
			CreateTime:    u.CreateTime.Unix(),
		})
	}

	return &user.GetUserListResponse{
		List:       usersList,
		TotalCount: totalCount,
		TotalPage:  int64(math.Ceil(float64(totalCount) / float64(in.PageSize))),
		CurrPage:   in.PageNumber,
		PageSize:   in.PageSize,
	}, nil
}
