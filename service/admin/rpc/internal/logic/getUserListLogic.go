package logic

import (
	"context"
	"errors"
	"math"

	"github.com/zeromicro/go-zero/core/logx"
	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"
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

func (l *GetUserListLogic) GetUserList(in *admin.GetUserListRequest) (*admin.GetUserListResponse, error) {
	// todo: add your logic here and delete this line
	// 分页
	limit := in.PageSize
	offset := in.PageSize * (in.PageNumber - 1)
	// 查询用户
	usersRes, err := l.svcCtx.UserModel.TakePart(l.ctx, limit, offset)
	if err != nil {
		logx.Error("用户列表获取失败" + err.Error())
		return nil, errors.New("用户列表获取失败" + err.Error())
	}
	usersList := make([]*admin.UserModel, 0)
	for _, user := range usersRes {
		usersList = append(usersList, &admin.UserModel{
			UserId:        user.UserId,
			NickName:      user.NickName,
			LoginName:     user.LoginName,
			PasswordMd5:   user.PasswordMd5,
			IntroduceSign: user.IntroduceSign,
			IsDeleted:     user.IsDeleted,
			LockedFlag:    user.LockedFlag,
			CreateTime:    user.CreateTime.Unix(),
		})
	}

	// 查询总数
	totalCount, err := l.svcCtx.UserModel.CountAll(l.ctx)
	if err != nil {
		logx.Error("用户总数获取失败" + err.Error())
		return nil, errors.New("用户总数获取失败" + err.Error())
	}

	return &admin.GetUserListResponse{
		List:       usersList,
		TotalCount: totalCount,
		TotalPage:  int64(math.Ceil(float64(totalCount) / float64(in.PageSize))),
		CurrPage:   in.PageNumber,
		PageSize:   in.PageSize,
	}, nil
}
