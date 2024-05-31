package handler

import (
	"github.com/gin-gonic/gin"
	"golang-crud-2024/core/entity"
	service "golang-crud-2024/core/service"
	"golang-crud-2024/pkg/api"
	"golang-crud-2024/pkg/context"
	"golang-crud-2024/pkg/fault"
	"net/http"
	"strconv"
)

type CartHandler interface {
	CreateCart(ctx *gin.Context)
	GetCart(ctx *gin.Context)
	GetCartByID(ctx *gin.Context)
	UpdateCart(ctx *gin.Context)
	DeleteCart(ctx *gin.Context)
}

type cartHandler struct {
	cartService service.CartService
}

func NewCart(cartService service.CartService) CartHandler {
	return &cartHandler{cartService: cartService}
}

func (c cartHandler) CreateCart(ctx *gin.Context) {
	var req *entity.CreateCart
	if err := ctx.ShouldBind(&req); err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPPreconditionFailedError, err.Error()))
		return
	}

	req.UserSerial = context.GetUserSerialFromGinContext(ctx)

	err := c.cartService.CreateCart(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, nil)
}

func (c cartHandler) GetCart(ctx *gin.Context) {
	userSerial := context.GetUserSerialFromGinContext(ctx)

	resp, err := c.cartService.GetCart(ctx, userSerial)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, resp)
}

func (c cartHandler) GetCartByID(ctx *gin.Context) {
	var req entity.GetCartByID
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	req.ID = int64(id)

	if err != nil {
		api.ResponseFailed(ctx, err)

	}
	req.UserSerial = context.GetUserSerialFromGinContext(ctx)

	resp, err := c.cartService.GetCartByID(ctx, &req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, resp)
}

func (c cartHandler) UpdateCart(ctx *gin.Context) {
	var req *entity.UpdateCart
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPPreconditionFailedError, err.Error()))
		return
	}

	req.UserSerial = context.GetUserSerialFromGinContext(ctx)

	err := c.cartService.UpdateCart(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, nil)
}

func (c cartHandler) DeleteCart(ctx *gin.Context) {
	var req *entity.DeleteCart
	if err := ctx.ShouldBind(&req); err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPPreconditionFailedError, err.Error()))
		return
	}

	req.UserSerial = context.GetUserSerialFromGinContext(ctx)

	err := c.cartService.DeleteCart(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, nil)
}
