package handler

import (
	"github.com/gin-gonic/gin"
	"golang-crud-2024/core/entity"
	service "golang-crud-2024/core/service"
	"golang-crud-2024/pkg/api"
	"golang-crud-2024/pkg/fault"
	"net/http"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

type authHandler struct {
	authService service.AuthService
}

func (a *authHandler) Register(ctx *gin.Context) {
	var req *entity.RegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPBadRequestError, err.Error()))
		return
	}

	err := a.authService.Register(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, nil)
}

func (a *authHandler) Login(ctx *gin.Context) {
	var req *entity.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.ResponseFailed(ctx, fault.ErrorDictionary(fault.HTTPBadRequestError, err.Error()))
		return
	}

	resp, err := a.authService.Login(ctx, req)
	if err != nil {
		api.ResponseFailed(ctx, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusOK, resp)
}
