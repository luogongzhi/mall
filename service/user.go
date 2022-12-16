package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/utils"
	"mall/serializer"
	"net/http"
	"time"
)

type UserService struct{}

// Register 用户注册
func (*UserService) Register(ctx context.Context, dto serializer.UserLoginRegisterDTO) serializer.ResponseResult {
	userDao := dao.NewUserDao(ctx)
	cartDao := dao.NewCartDao(ctx)

	// 判断用户名是否存在
	_, exist, err := userDao.ExistOrNotByUserName(dto.Username)
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
	id, err := userDao.Create(&model.User{
		Username: dto.Username,
		Password: utils.MD5(dto.Password),
	})
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 初始化用户购物车
	err = cartDao.Create(&model.Cart{
		UserId: id,
		Total:  0,
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
func (*UserService) Login(ctx context.Context, dto serializer.UserLoginRegisterDTO) serializer.ResponseResult {
	userDao := dao.NewUserDao(ctx)

	// 判断用户名是否存在
	user, exist, err := userDao.ExistOrNotByUserName(dto.Username)
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
	if utils.MD5(dto.Password) != user.Password {
		return serializer.ResponseResult{
			Code: e.ErrorNotCompare,
			Msg:  e.GetMsg(e.ErrorNotCompare),
		}
	}

	// 生成token
	var token string
	if token, err = utils.GenerateToken(user.Id, dto.Username); err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorAuthToken,
			Msg:  e.GetMsg(e.ErrorAuthToken),
		}
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
		Data: map[string]interface{}{
			"user":  serializer.NewUserVO(user),
			"token": token,
		},
	}
}

// Detail 根据Id查询用户信息
func (*UserService) Detail(ctx context.Context, id uint64) serializer.ResponseResult {
	userDao := dao.NewUserDao(ctx)

	// 根据id查询用户
	user, _, _ := userDao.GetById(id)
	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
		Data: map[string]interface{}{
			"user": serializer.NewUserVO(user),
		},
	}
}

// Update 修改用户信息
func (*UserService) Update(ctx context.Context, dto serializer.UserUpdateDTO, id uint64) serializer.ResponseResult {
	userDao := dao.NewUserDao(ctx)

	var gender uint
	switch dto.Gender {
	case "男":
		gender = 1
	case "女":
		gender = 2
	default:
		gender = 3
	}

	birth, err := time.ParseInLocation("2006-01-02", dto.Birth, time.Local)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDate,
			Msg:  e.GetMsg(e.ErrorDate),
		}
	}

	// 根据Id修改用户信息
	err = userDao.UpdateById(id, &model.User{
		Username: dto.Username,
		Tel:      dto.Tel,
		Email:    dto.Email,
		Gender:   gender,
		Birth:    birth,
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
