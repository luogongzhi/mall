package service

import (
	"context"
	"mall/dao"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
)

type CartService struct{}

func (service *CartService) Detail(ctx context.Context, userId uint64) serializer.ResponseResult {
	cartDao := dao.NewCartDao(ctx)
	cartProductDao := dao.NewCartProductDao(ctx)

	cart, err := cartDao.GetCartByUserId(userId)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	cartProductList, err := cartProductDao.GetCartProductList(cart.Id)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
		Data: map[string]interface{}{
			"cart": serializer.NewCartVO(cart, *cartProductList),
		},
	}
}
