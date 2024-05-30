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
	productHandler := handler.NewProduct(app.ProductService)
	cartHandler := handler.NewCart(app.CartService)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/register", authHandler.Register)

	router.Use(auth.AuthMiddleware(app.Cfg)).GET("/product-category", productCategoryHandler.GetProductCategory)
	router.Use(auth.AuthMiddleware(app.Cfg)).GET("/product", productHandler.GetProduct)

	router.Use(auth.AuthMiddleware(app.Cfg)).POST("/cart", cartHandler.CreateCart)
	router.Use(auth.AuthMiddleware(app.Cfg)).GET("/cart", cartHandler.GetCart)
	router.Use(auth.AuthMiddleware(app.Cfg)).GET("/cart/:id", cartHandler.GetCartByID)
	router.Use(auth.AuthMiddleware(app.Cfg)).PUT("/cart", cartHandler.UpdateCart)
	router.Use(auth.AuthMiddleware(app.Cfg)).DELETE("/cart", cartHandler.DeleteCart)

	return router
}
