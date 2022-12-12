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

// Register 用户注册
func (service UserService) Register(ctx context.Context) serializer.ResponseResult {

	userDao := dao.NewUserDao(ctx)
	// 判断用户名是否存在
	_, exist, err := userDao.ExistOrNotByUserName(service.Username)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	if exist {
		return serializer.ResponseResult{
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
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}

// Login 用户登录
func (service UserService) Login(ctx context.Context) serializer.ResponseResult {

	userDao := dao.NewUserDao(ctx)
	// 判断用户名是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.Username)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	if !exist {
		return serializer.ResponseResult{
			Code: e.ErrorNotExistUser,
			Msg:  e.GetMsg(e.ErrorNotExistUser),
		}
	}

	// 校验密码
	if utils.MD5(service.Password) != user.Password {
		return serializer.ResponseResult{
			Code: e.ErrorNotCompare,
			Msg:  e.GetMsg(e.ErrorNotCompare),
		}
	}

	//生成token
	token, err := utils.GenerateToken(user.Id, service.Username)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorAuthToken,
			Msg:  e.GetMsg(e.ErrorAuthToken),
		}
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
		Data: map[string]interface{}{
			"user":  serializer.BuildUserVO(user),
			"token": token,
		},
	}
}
