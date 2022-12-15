package service

import (
	"context"
	"github.com/shopspring/decimal"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
)

type OrderService struct{}

// Create 创建订单
func (service *OrderService) Create(ctx context.Context, userId uint64, dto serializer.OrderCreateDTO) serializer.ResponseResult {
	cartDao := dao.NewCartDao(ctx)
	cartProductDao := dao.NewCartProductDao(ctx)
	productDao := dao.NewProductDao(ctx)
	orderDao := dao.NewOrderDao(ctx)
	orderProductDao := dao.NewOrderProductDao(ctx)

	// 1.清空购物车
	// 根据用户查询该用户购物车
	cart, err := cartDao.GetByUserId(userId)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	// 判断购物车商品是否为空
	if cart.Total == 0 {
		return serializer.ResponseResult{
			Code: e.ErrorNotExistCartProduct,
			Msg:  e.GetMsg(e.ErrorNotExistCartProduct),
		}
	}
	// 根据用户购物车Id查询购物车中的商品
	var cartProductList *[]serializer.CartProductVO
	cartProductList, err = cartProductDao.GetList(cart.Id)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	// 检查商品库存
	var productList []model.Product
	for _, cartProduct := range *cartProductList {
		var exist bool
		var product *model.Product
		product, exist, err = productDao.GetById(cartProduct.ProductId)
		if err != nil {
			return serializer.ResponseResult{
				Code: e.ErrorDatabase,
				Msg:  e.GetMsg(e.ErrorDatabase),
			}
		}
		// 判断商品是否存在
		if !exist {
			return serializer.ResponseResult{
				Code: e.ErrorNotExistProduct,
				Msg:  e.GetMsg(e.ErrorNotExistProduct),
				Data: cartProduct.ProductId,
			}
		}
		// 如果购物车中的商品件数 > 该商品库存
		if cartProduct.Total > product.Total {
			return serializer.ResponseResult{
				Code: e.ErrorNotEnoughProduct,
				Msg:  e.GetMsg(e.ErrorNotEnoughProduct),
				Data: cartProduct.ProductId,
			}
		}
		productList = append(productList, *product)
	}
	// 根据用户购物车Id删除购物车中的商品
	err = cartProductDao.DeleteByCartId(cart.Id)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	// 修改用户购物车商品总数
	err = cartDao.UpdateTotal(&model.Cart{
		UserId: userId,
		Total:  0,
		Model: model.Model{
			CreatedAt: cart.Model.CreatedAt,
		},
	})
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 2.修改商品库存
	// 商品总金额
	var productAmount float64
	for i, cartProduct := range *cartProductList {
		err = productDao.UpdateProductTotal(&model.Product{
			Model: model.Model{Id: cartProduct.ProductId},
			Total: productList[i].Total - cartProduct.Total,
		})
		if err != nil {
			return serializer.ResponseResult{
				Code: e.ErrorDatabase,
				Msg:  e.GetMsg(e.ErrorDatabase),
			}
		}
		num := decimal.NewFromFloat(cartProduct.Price).Mul(decimal.NewFromFloat(float64(cartProduct.Total)))
		f, _ := num.Float64()
		productAmount += f
	}

	// 3.创建订单
	// 创建 order
	var orderId uint64
	orderId, err = orderDao.Create(&model.Order{
		UserId:         userId,
		AddressName:    dto.AddressName,
		AddressTel:     dto.AddressTel,
		AddressDetails: dto.AddressDetails,
		ProductAmount:  productAmount,
		FreightAmount:  6,
		TotalAmount:    productAmount + 6,
		Status:         false,
	})
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 创建 order_product
	for _, cartProduct := range *cartProductList {
		err = orderProductDao.Create(&model.OrderProduct{
			OrderId:   orderId,
			ProductId: cartProduct.ProductId,
			Total:     cartProduct.Total,
		})
		if err != nil {
			return serializer.ResponseResult{
				Code: e.ErrorDatabase,
				Msg:  e.GetMsg(e.ErrorDatabase),
			}
		}
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}
