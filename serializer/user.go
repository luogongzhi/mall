package serializer

import (
	"mall/model"
)

type UserLoginRegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateDTO struct {
	Username string `json:"username"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Birth    string `json:"birth"`
}

type UserVO struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Birth    string `json:"birth"`
}

func NewUserVO(user *model.User) UserVO {
	var gender string
	switch user.Gender {
	case 1:
		gender = "男"
	case 2:
		gender = "女"
	default:
		gender = "未知"
	}

	return UserVO{
		Id:       user.Id,
		Username: user.Username,
		Tel:      user.Tel,
		Email:    user.Email,
		Gender:   gender,
		Birth:    user.Birth.Format("2006-01-02"),
	}
}
