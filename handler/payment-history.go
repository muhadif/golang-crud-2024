package handler

import (
	"github.com/gin-gonic/gin"
	"golang-crud-2024/core/entity"
	service "golang-crud-2024/core/service"
	"golang-crud-2024/pkg/api"
	"golang-crud-2024/pkg/context"
	"net/http"
)

type PaymentHistoryHandler interface {
	CreatePayment(ctx *gin.Context)
	GetPaymentHistory(ctx *gin.Context)
}

type paymentHistoryHandler struct {
	paymentHistory service.PaymentHistoryService
}

func NewPaymentHistory(paymentHistory service.PaymentHistoryService) PaymentHistoryHandler {
	return &paymentHistoryHandler{paymentHistory: paymentHistory}
}

func (p paymentHistoryHandler) CreatePayment(ctx *gin.Context) {
	var req *entity.CreatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	req.UserSerial = context.GetUserSerialFromGinContext(ctx)

	err := p.paymentHistory.CreatePayment(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, nil)
}

func (p paymentHistoryHandler) GetPaymentHistory(ctx *gin.Context) {
	userSerial := context.GetUserSerialFromGinContext(ctx)
	resp, err := p.paymentHistory.GetPaymentHistory(ctx, userSerial)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, resp)
}
