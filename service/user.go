package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/utils"
	"mall/serializer"
	"net/http"
)

type UserService struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (service UserService) Register(ctx context.Context) serializer.Response {

	userDao := dao.NewUserDao(ctx)
	// 判断用户名是否存在
	exist, err := userDao.ExistOrNotByUserName(service.Username)
	if err != nil {
		return serializer.Response{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	if exist {
		return serializer.Response{
			Code: e.ErrorExistUser,
			Msg:  e.GetMsg(e.ErrorExistUser),
		}
	}

	// 创建用户
	err = userDao.CreateUser(&model.User{
		Username: service.Username,
		Password: utils.MD5(service.Password),
	})
	if err != nil {
		return serializer.Response{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	return serializer.Response{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}

}
