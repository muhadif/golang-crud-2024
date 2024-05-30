package handler

import (
	"github.com/gin-gonic/gin"
	"golang-crud-2024/core/entity"
	service "golang-crud-2024/core/service"
	"golang-crud-2024/pkg/api"
	"golang-crud-2024/pkg/context"
	"net/http"
)

type CheckoutHandler interface {
	CreateCheckout(ctx *gin.Context)
	GetCurrentCheckout(ctx *gin.Context)
}

type checkoutHandler struct {
	checkoutService service.CheckoutService
}

func NewCheckoutHandler(checkoutService service.CheckoutService) CheckoutHandler {
	return &checkoutHandler{
		checkoutService: checkoutService,
	}
}

func (c checkoutHandler) CreateCheckout(ctx *gin.Context) {
	var req *entity.CreateCheckoutSession
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	req.UserSerial = context.GetUserSerialFromGinContext(ctx)

	err := c.checkoutService.CreateCheckout(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, nil)
}

func (c checkoutHandler) GetCurrentCheckout(ctx *gin.Context) {
	userSerial := context.GetUserSerialFromGinContext(ctx)

	resp, err := c.checkoutService.GetCurrentCheckout(ctx, userSerial)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, resp)
}
