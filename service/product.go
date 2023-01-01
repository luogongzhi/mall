package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
)

type ProductService struct{}

func (*ProductService) Detail(ctx context.Context, id uint64) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	productDao := dao.NewProductDao(db)
	product, exist, err := productDao.GetById(id)
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

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
		Data: map[string]interface{}{
			"product": serializer.NewProductVO(product),
		},
	}
}

func (*ProductService) Create(ctx context.Context, dto serializer.ProductCreateDTO) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	productDao := dao.NewProductDao(db)
	err := productDao.Create(&model.Product{
		Title:     dto.Title,
		Info:      dto.Info,
		AttrValue: dto.AttrValue,
		Price:     dto.Price,
		Total:     dto.Total,
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

func (*ProductService) Delete(ctx context.Context, id uint64) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	productDao := dao.NewProductDao(db)
	_, exist, err := productDao.GetById(id)
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

	err = productDao.Delete(id)
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

func (*ProductService) Update(ctx context.Context, dto serializer.ProductUpdateDTO) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	productDao := dao.NewProductDao(db)
	_, exist, err := productDao.GetById(dto.Id)
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

	err = productDao.UpdateProduct(&model.Product{
		Model:     model.Model{Id: dto.Id},
		Title:     dto.Title,
		Info:      dto.Info,
		AttrValue: dto.AttrValue,
		Price:     dto.Price,
		Total:     dto.Total,
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

func (*ProductService) List(ctx context.Context, dto serializer.PaginateDTO) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	productDao := dao.NewProductDao(db)

	pageNum := dto.PageNum
	pageSize := dto.PageSize
	if pageNum <= 0 {
		pageNum = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	productList, err := productDao.GetList(pageNum, pageSize)
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
			"productList": serializer.NewProductVOList(productList),
		},
	}
}
