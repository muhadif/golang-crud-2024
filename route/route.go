package route

import (
	"github.com/gin-gonic/gin"
	"golang-crud-2024/app"
	"golang-crud-2024/handler"
	"golang-crud-2024/pkg/auth"
)

func NewRouter(app *app.App) *gin.Engine {
	authHandler := handler.NewAuthHandler(app.AuthService)
	productCategoryHandler := handler.NewProductCategory(app.ProductCategoryService)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(auth.AuthMiddleware(true))

	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/register", authHandler.Register)

	router.Use(auth.AuthMiddleware(app.Cfg)).GET("product-category", productCategoryHandler.GetProductCategory)

	return router
}
