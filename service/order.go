package service

import (
	"context"
	"github.com/shopspring/decimal"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
	"sync"
)

type OrderService struct{}

// putResponseResultChan 将返回结果放到通道
func putResponseResultChan(code int, data map[string]uint64, responseResultChan *chan serializer.ResponseResult) {
	result := serializer.ResponseResult{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
	*responseResultChan <- result
}

// cleanCart 清空购物车
func cleanCart(wg *sync.WaitGroup, responseResultChan *chan serializer.ResponseResult, productListChan *chan []model.Product, userId uint64, cartId uint64, cartProductList *[]serializer.CartProductVO, cartDao *dao.CartDao, cartProductDao *dao.CartProductDao, productDao *dao.ProductDao) {
	defer wg.Done()

	// 检查商品库存
	var productList []model.Product
	for _, cartProduct := range *cartProductList {
		product, exist, err := productDao.GetById(cartProduct.ProductId)
		if err != nil {
			putResponseResultChan(e.ErrorDatabase, nil, responseResultChan)
			close(*productListChan)
			return
		}
		// 判断商品是否存在
		if !exist {
			putResponseResultChan(e.ErrorNotExistProduct, map[string]uint64{"productId": cartProduct.ProductId}, responseResultChan)
			close(*productListChan)
			return
		}
		// 如果购物车中的商品件数 > 该商品库存
		if cartProduct.Total > product.Total {
			putResponseResultChan(e.ErrorNotEnoughProduct, map[string]uint64{"productId": cartProduct.ProductId}, responseResultChan)
			close(*productListChan)
			return
		}
		productList = append(productList, *product)
	}
	*productListChan <- productList

	// 根据用户购物车Id删除购物车中的商品
	if err := cartProductDao.DeleteByCartId(cartId); err != nil {
		putResponseResultChan(e.ErrorDatabase, nil, responseResultChan)
		return
	}

	// 修改用户购物车商品总数
	err := cartDao.UpdateTotal(&model.Cart{
		UserId: userId,
		Total:  0,
	})
	if err != nil {
		putResponseResultChan(e.ErrorDatabase, nil, responseResultChan)
		return
	}
}

// updateProductTotal 修改商品库存
func updateProductTotal(wg *sync.WaitGroup, responseResultChan *chan serializer.ResponseResult, productListChan *chan []model.Product, productDao *dao.ProductDao, cartProductList *[]serializer.CartProductVO) {
	defer wg.Done()

	// 如果chan已关闭就返回
	if productList, ok := <-*productListChan; !ok {
		return
	} else {
		for i, cartProduct := range *cartProductList {
			err := productDao.UpdateProductTotal(&model.Product{
				Model: model.Model{Id: cartProduct.ProductId},
				Total: productList[i].Total - cartProduct.Total,
			})
			if err != nil {
				putResponseResultChan(e.ErrorDatabase, nil, responseResultChan)
				return
			}
		}
	}
}

// createOrder 创建订单
func createOrder(wg *sync.WaitGroup, responseResultChan *chan serializer.ResponseResult, cartProductList *[]serializer.CartProductVO, userId uint64, dto *serializer.OrderCreateDTO, orderDao *dao.OrderDao, orderProductDao *dao.OrderProductDao) {
	defer wg.Done()

	var productAmount float64
	for _, cartProduct := range *cartProductList {
		f, _ := decimal.NewFromFloat(cartProduct.Price).Mul(decimal.NewFromFloat(float64(cartProduct.Total))).Float64()
		productAmount += f
	}

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
		putResponseResultChan(e.ErrorDatabase, nil, responseResultChan)
		return
	}

	// 创建 order_product
	for _, cartProduct := range *cartProductList {
		err = orderProductDao.Create(&model.OrderProduct{
			OrderId:   orderId,
			ProductId: cartProduct.ProductId,
			Total:     cartProduct.Total,
		})
		if err != nil {
			putResponseResultChan(e.ErrorDatabase, nil, responseResultChan)
			return
		}
	}
}

func (*OrderService) Create(ctx context.Context, userId uint64, dto serializer.OrderCreateDTO) serializer.ResponseResult {
	var wg sync.WaitGroup
	wg.Add(3)

	responseResultChan := make(chan serializer.ResponseResult, 3)
	productListChan := make(chan []model.Product)

	db := dao.NewTransactionDBClient(ctx)
	cartDao := dao.NewCartDao(db)
	cartProductDao := dao.NewCartProductDao(db)
	productDao := dao.NewProductDao(db)
	orderDao := dao.NewOrderDao(db)
	orderProductDao := dao.NewOrderProductDao(db)

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
	cartProductList, err := cartProductDao.GetList(cart.Id)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	// 1. 清空购物车
	go cleanCart(&wg, &responseResultChan, &productListChan, userId, cart.Id, cartProductList, cartDao, cartProductDao, productDao)
	// 2. 修改商品库存
	go updateProductTotal(&wg, &responseResultChan, &productListChan, productDao, cartProductList)
	// 3.创建订单
	go createOrder(&wg, &responseResultChan, cartProductList, userId, &dto, orderDao, orderProductDao)

	wg.Wait()
	close(responseResultChan)
	// 如果从chan拿到结果返回就代表有异常，回滚，返回
	if responseResult, ok := <-responseResultChan; ok {
		db.Rollback()
		return responseResult
	}
	db.Commit()

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}

func (*OrderService) Update(ctx context.Context, userId uint64, dto serializer.OrderUpdateDTO) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	orderDao := dao.NewOrderDao(db)

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
	db := dao.NewDBClient(ctx)
	orderDao := dao.NewOrderDao(db)
	orderProductDao := dao.NewOrderProductDao(db)

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
	db := dao.NewTransactionDBClient(ctx)
	orderDao := dao.NewOrderDao(db)
	orderProductDao := dao.NewOrderProductDao(db)
	productDao := dao.NewProductDao(db)

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
		db.Rollback()
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
			db.Rollback()
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
				db.Rollback()
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
		db.Rollback()
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	db.Commit()
	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}
