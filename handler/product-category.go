package handler

import (
	"github.com/gin-gonic/gin"
	service "golang-crud-2024/core/service"
	"golang-crud-2024/pkg/api"
	"net/http"
)

type ProductCategoryHandler interface {
	GetProductCategory(ctx *gin.Context)
}

type productCategoryHandler struct {
	productCategoryService service.ProductCategoryService
}

func NewProductCategory(productCategoryService service.ProductCategoryService) ProductCategoryHandler {
	return &productCategoryHandler{productCategoryService: productCategoryService}
}

func (p productCategoryHandler) GetProductCategory(ctx *gin.Context) {
	resp, err := p.productCategoryService.GetProductCategory(ctx)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, resp)
}
