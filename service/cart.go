package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
)

type CartService struct{}

func (service *CartService) Detail(ctx context.Context, userId uint64) serializer.ResponseResult {
	cartDao := dao.NewCartDao(ctx)
	cartProductDao := dao.NewCartProductDao(ctx)

	cart, err := cartDao.GetByUserId(userId)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	cartProductList, err := cartProductDao.GetList(cart.Id)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	if cartProductList == nil {
		return serializer.ResponseResult{
			Code: http.StatusOK,
			Msg:  e.GetMsg(http.StatusOK),
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

func (service *CartService) Create(ctx context.Context, dto serializer.CartCreateDeleteDTO, id uint64) serializer.ResponseResult {
	cartDao := dao.NewCartDao(ctx)
	cartProductDao := dao.NewCartProductDao(ctx)

	// 根据用户查询该用户购物车
	cart, err := cartDao.GetByUserId(id)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 查询该商品是否存在
	cartProduct, exist, err := cartProductDao.GetByProductId(&model.CartProduct{
		CartId:    cart.Id,
		ProductId: dto.ProductId,
	})
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 如果存在则修改商品件数，不存在则新建
	if exist {
		err = cartProductDao.UpdateTotalByProductId(&model.CartProduct{
			CartId:    cart.Id,
			ProductId: dto.ProductId,
			Total:     cartProduct.Total + dto.Total,
		})
		if err != nil {
			return serializer.ResponseResult{
				Code: e.ErrorDatabase,
				Msg:  e.GetMsg(e.ErrorDatabase),
			}
		}
	} else {
		err = cartProductDao.Create(&model.CartProduct{
			CartId:    cart.Id,
			ProductId: dto.ProductId,
			Total:     dto.Total,
		})
		if err != nil {
			return serializer.ResponseResult{
				Code: e.ErrorDatabase,
				Msg:  e.GetMsg(e.ErrorDatabase),
			}
		}
	}

	// 修改购物车商品总数
	err = cartDao.UpdateTotal(&model.Cart{
		UserId: id,
		Total:  cart.Total + dto.Total,
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

func (service *CartService) Delete(ctx context.Context, dto serializer.CartCreateDeleteDTO, id uint64) serializer.ResponseResult {
	cartDao := dao.NewCartDao(ctx)
	cartProductDao := dao.NewCartProductDao(ctx)

	// 根据用户查询该用户购物车
	cart, err := cartDao.GetByUserId(id)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 查询该商品是否存在
	cartProduct, exist, err := cartProductDao.GetByProductId(&model.CartProduct{
		CartId:    cart.Id,
		ProductId: dto.ProductId,
	})
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	if !exist {
		return serializer.ResponseResult{
			Code: e.ErrorNotExistProduct,
			Msg:  e.GetMsg(e.ErrorNotExistProduct),
		}
	}

	// 如果传入 > 就报错； = 就删除购物车中商品；< 就修改商品件数
	if dto.Total > cartProduct.Total {
		return serializer.ResponseResult{
			Code: e.ErrorProductTotal,
			Msg:  e.GetMsg(e.ErrorProductTotal),
		}
	} else if dto.Total == cartProduct.Total {
		err = cartProductDao.DeleteByProductId(&model.CartProduct{
			CartId:    cart.Id,
			ProductId: dto.ProductId,
		})
	} else {
		err = cartProductDao.UpdateTotalByProductId(&model.CartProduct{
			CartId:    cart.Id,
			ProductId: dto.ProductId,
			Total:     cartProduct.Total - dto.Total,
		})
	}
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 修改购物车商品总数
	err = cartDao.UpdateTotal(&model.Cart{
		UserId: id,
		Total:  cart.Total - dto.Total,
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
