package handler

import (
	"github.com/gin-gonic/gin"
	"golang-crud-2024/core/entity"
	service "golang-crud-2024/core/service"
	"golang-crud-2024/pkg/api"
	"golang-crud-2024/pkg/fault"
	"net/http"
)

type PaymentCallbackHandler interface {
	CallbackPaymentVATransfer(ctx *gin.Context)
}

type paymentCallbackHandler struct {
	paymentCallbackService service.PaymentCallbackService
}

func NewPaymentCallbackHandler(paymentCallbackService service.PaymentCallbackService) PaymentCallbackHandler {
	return &paymentCallbackHandler{paymentCallbackService: paymentCallbackService}
}

func (p paymentCallbackHandler) CallbackPaymentVATransfer(ctx *gin.Context) {
	var req *entity.PaymentCallbackVATransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPPreconditionFailedError, err.Error()))
		return
	}

	err := p.paymentCallbackService.CallbackPaymentVATransfer(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPPreconditionFailedError, err.Error()))
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, nil)
}
