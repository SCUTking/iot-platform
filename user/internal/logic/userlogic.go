package logic

import (
	"context"
	"errors"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/user/internal/svc"
	"iot-platform/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserLoginRequest) (resp *types.UserLoginReply, err error) {
	// todo: add your logic here and delete this line
	var user models.UserBasic
	err = l.svcCtx.DB.Where(&models.UserBasic{Name: req.Username, Password: helper.Md5(req.Password)}).Find(&user).Error
	if err != nil {
		logx.Error("[DB error] : ", err)
		err = errors.New("用户名或者密码不正确")
		return
	}
	token, err := helper.GenerateToken(user.ID, user.Identity, user.Name, 3600*24*30)
	if err != nil {
		logx.Error("[DB error] : ", err)
		err = errors.New("用户名或者密码不正确")
		return
	}
	resp = &types.UserLoginReply{Token: token}
	return
}
