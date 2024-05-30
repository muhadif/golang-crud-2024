package context

import "github.com/gin-gonic/gin"

func GetUserSerialFromGinContext(ctx *gin.Context) string {
	userSerial, _ := ctx.Get("userSerial")
	return userSerial.(string)
}
