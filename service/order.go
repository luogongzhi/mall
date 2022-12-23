package service

import (
	"context"
	"github.com/shopspring/decimal"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/utils"
	"mall/serializer"
	"net/http"
)

type OrderService struct{}

// cleanCart 清空购物车
func cleanCart(userId uint64, cartDao *dao.CartDao, cartProductDao *dao.CartProductDao, productDao *dao.ProductDao) (code int, data map[string]uint64, cartProductList *[]serializer.CartProductVO, productList []model.Product) {
	// 根据用户查询该用户购物车
	cart, err := cartDao.GetByUserId(userId)
	if err != nil {
		return e.ErrorDatabase, nil, nil, nil
	}
	// 判断购物车商品是否为空
	if cart.Total == 0 {
		return e.ErrorNotExistCartProduct, nil, nil, nil
	}
	// 根据用户购物车Id查询购物车中的商品
	cartProductList, err = cartProductDao.GetList(cart.Id)
	if err != nil {
		return e.ErrorDatabase, nil, nil, nil
	}
	// 检查商品库存
	for _, cartProduct := range *cartProductList {
		product, exist, err := productDao.GetById(cartProduct.ProductId)
		if err != nil {
			return e.ErrorDatabase, nil, nil, nil
		}
		// 判断商品是否存在
		if !exist {
			return e.ErrorNotExistProduct, map[string]uint64{"productId": cartProduct.ProductId}, nil, nil
		}
		// 如果购物车中的商品件数 > 该商品库存
		if cartProduct.Total > product.Total {
			return e.ErrorNotEnoughProduct, map[string]uint64{"productId": cartProduct.ProductId}, nil, nil
		}
		productList = append(productList, *product)
	}
	// 根据用户购物车Id删除购物车中的商品
	err = cartProductDao.DeleteByCartId(cart.Id)
	if err != nil {
		return e.ErrorDatabase, nil, nil, nil
	}
	// 修改用户购物车商品总数
	err = cartDao.UpdateTotal(&model.Cart{
		UserId: userId,
		Total:  0,
	})
	if err != nil {
		return e.ErrorDatabase, nil, nil, nil
	}
	return 0, nil, cartProductList, productList
}

// updateProductTotal 修改商品库存
func updateProductTotal(productDao *dao.ProductDao, cartProductList *[]serializer.CartProductVO, productList []model.Product) (code int, productAmount float64) {
	for i, cartProduct := range *cartProductList {
		err := productDao.UpdateProductTotal(&model.Product{
			Model: model.Model{Id: cartProduct.ProductId},
			Total: productList[i].Total - cartProduct.Total,
		})
		if err != nil {
			return e.ErrorDatabase, 0
		}
		f, _ := decimal.NewFromFloat(cartProduct.Price).Mul(decimal.NewFromFloat(float64(cartProduct.Total))).Float64()
		productAmount += f
	}
	return 0, productAmount
}

// createOrder 创建订单
func createOrder(userId uint64, productAmount float64, dto *serializer.OrderCreateDTO, cartProductList *[]serializer.CartProductVO, orderDao *dao.OrderDao, orderProductDao *dao.OrderProductDao) (code int) {
	// 创建 order
	orderId, err := orderDao.Create(&model.Order{
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
		return e.ErrorDatabase
	}
	// 创建 order_product
	for _, cartProduct := range *cartProductList {
		err = orderProductDao.Create(&model.OrderProduct{
			OrderId:   orderId,
			ProductId: cartProduct.ProductId,
			Total:     cartProduct.Total,
		})
		if err != nil {
			return e.ErrorDatabase
		}
	}
	return 0
}

func (*OrderService) Create(ctx context.Context, userId uint64, dto serializer.OrderCreateDTO) serializer.ResponseResult {
	cartDao := dao.NewCartTransactionDao(ctx)
	cartProductDao := dao.NewCartProductTransactionDao(ctx)
	productDao := dao.NewProductTransactionDao(ctx)
	orderDao := dao.NewOrderTransactionDao(ctx)
	orderProductDao := dao.NewOrderProductTransactionDao(ctx)

	// 1. 清空购物车
	code, data, cartProductList, productList := cleanCart(userId, cartDao, cartProductDao, productDao)
	if code != 0 {
		utils.Rollback(cartDao, cartProductDao, productDao, orderDao, orderProductDao)
		return serializer.ResponseResult{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: data,
		}
	}
	// 2.修改商品库存
	code, productAmount := updateProductTotal(productDao, cartProductList, productList)
	if code != 0 {
		utils.Rollback(cartDao, cartProductDao, productDao, orderDao, orderProductDao)
		return serializer.ResponseResult{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	// 3.创建订单
	code = createOrder(userId, productAmount, &dto, cartProductList, orderDao, orderProductDao)
	if code != 0 {
		utils.Rollback(cartDao, cartProductDao, productDao, orderDao, orderProductDao)
		return serializer.ResponseResult{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}

	utils.Commit(cartDao, cartProductDao, productDao, orderDao, orderProductDao)
	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}

func (*OrderService) Update(ctx context.Context, userId uint64, dto serializer.OrderUpdateDTO) serializer.ResponseResult {
	orderDao := dao.NewOrderDao(ctx)

	//判断是否存在该订单Id
	_, exist, err := orderDao.GetByOrderId(userId, dto.OrderId)
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

	// 修改
	status := false
	if dto.Status == "已完成" {
		status = true
	}
	err = orderDao.Update(&model.Order{
		Model:          model.Model{Id: dto.OrderId},
		AddressName:    dto.AddressName,
		AddressTel:     dto.AddressTel,
		AddressDetails: dto.AddressDetails,
		Status:         status,
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

func (*OrderService) List(ctx context.Context, userId uint64) serializer.ResponseResult {
	orderDao := dao.NewOrderDao(ctx)
	orderProductDao := dao.NewOrderProductDao(ctx)

	// 获取用户订单列
	orderList, err := orderDao.GetListByUserId(userId)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	if orderList == nil {
		return serializer.ResponseResult{
			Code: http.StatusOK,
			Msg:  e.GetMsg(http.StatusOK),
		}
	}

	// 封装返回的订单列
	var orderVOList []*serializer.OrderVO
	for _, order := range *orderList {
		var orderProductVOList []serializer.OrderProductVO
		var orderProductList *[]model.OrderProduct
		// 获得对于订单的订单商品列
		orderProductList, err = orderProductDao.GetListByOrderId(order.Id)
		if err != nil {
			return serializer.ResponseResult{
				Code: e.ErrorDatabase,
				Msg:  e.GetMsg(e.ErrorDatabase),
			}
		}
		for _, orderProduct := range *orderProductList {
			orderProductVO := serializer.OrderProductVO{
				ProductId: orderProduct.ProductId,
				Total:     orderProduct.Total,
			}
			orderProductVOList = append(orderProductVOList, orderProductVO)
		}
		orderVO := serializer.NewOrderVO(&order, &orderProductVOList)
		orderVOList = append(orderVOList, &orderVO)
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
		Data: orderVOList,
	}
}

func (*OrderService) Delete(ctx context.Context, userId uint64, orderId uint64) serializer.ResponseResult {
	orderDao := dao.NewOrderDao(ctx)
	orderProductDao := dao.NewOrderProductDao(ctx)
	productDao := dao.NewProductDao(ctx)

	//订单是不是用户的
	order, exist, err := orderDao.GetByOrderId(userId, orderId)
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

	// 订单是否未完成
	if order.Status {
		return serializer.ResponseResult{
			Code: e.ErrortOrderFinished,
			Msg:  e.GetMsg(e.ErrortOrderFinished),
		}
	}

	// 查询订单的order_product
	var orderProductList *[]model.OrderProduct
	orderProductList, err = orderProductDao.GetListByOrderId(orderId)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 根据订单Id删除order_product
	err = orderProductDao.DeleteByOrderId(orderId)
	if err != nil {
		utils.Rollback(productDao, orderDao, orderProductDao)
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 遍历 order_product list，根据product_id查询对应商品
	for _, orderProduct := range *orderProductList {
		var product *model.Product
		product, exist, err = productDao.GetById(orderProduct.ProductId)
		if err != nil {
			utils.Rollback(productDao, orderDao, orderProductDao)
			return serializer.ResponseResult{
				Code: e.ErrorDatabase,
				Msg:  e.GetMsg(e.ErrorDatabase),
			}
		}
		//如果商品存在，则将商品total加上
		if exist {
			err = productDao.UpdateProductTotal(&model.Product{
				Model: model.Model{Id: orderProduct.ProductId},
				Total: product.Total + orderProduct.Total,
			})
			if err != nil {
				utils.Rollback(productDao, orderDao, orderProductDao)
				return serializer.ResponseResult{
					Code: e.ErrorDatabase,
					Msg:  e.GetMsg(e.ErrorDatabase),
				}
			}
		}
	}

	// 删除order
	err = orderDao.DeleteById(orderId)
	if err != nil {
		utils.Rollback(productDao, orderDao, orderProductDao)
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	utils.Commit(productDao, orderDao, orderProductDao)
	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}
