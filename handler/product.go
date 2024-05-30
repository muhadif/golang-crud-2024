package handler

import (
	"github.com/gin-gonic/gin"
	service "golang-crud-2024/core/service"
	"golang-crud-2024/pkg/api"
	"net/http"
)

type ProductHandler interface {
	GetProduct(ctx *gin.Context)
}

type productHandler struct {
	productService service.ProductService
}

func NewProduct(productService service.ProductService) ProductHandler {
	return &productHandler{productService: productService}
}

func (p productHandler) GetProduct(ctx *gin.Context) {
	resp, err := p.productService.GetProduct(ctx)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, resp)
}
