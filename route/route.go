package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-limiter/memorystore"
	"golang-crud-2024/app"
	"golang-crud-2024/handler"
	"golang-crud-2024/pkg/security"
	"time"
)

func NewRouter(app *app.App) *gin.Engine {
	authHandler := handler.NewAuthHandler(app.AuthService)
	productCategoryHandler := handler.NewProductCategory(app.ProductCategoryService)
	productHandler := handler.NewProduct(app.ProductService)
	cartHandler := handler.NewCart(app.CartService)
	checkoutHandler := handler.NewCheckoutHandler(app.CheckoutService)
	paymentHistoryHandler := handler.NewPaymentHistory(app.PaymentHistoryService)
	paymentCallbackHandler := handler.NewPaymentCallbackHandler(app.PaymentCallbackService)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	store, err := memorystore.New(&memorystore.Config{
		Tokens:   50,
		Interval: time.Minute,
	})
	if err != nil {
		panic(err)
	}

	router.Use(security.RateLimiterMiddleware(store))

	// auth pth
	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/register", authHandler.Register)

	// callback payment path
	router.POST("/callback/payment/va-transfer", paymentCallbackHandler.CallbackPaymentVATransfer)

	// customer path
	router.Use(app.AuthMiddleware.AuthMiddleware()).GET("/product-category", productCategoryHandler.GetProductCategory)
	router.Use(app.AuthMiddleware.AuthMiddleware()).GET("/product", productHandler.GetProduct)

	router.Use(app.AuthMiddleware.AuthMiddleware()).POST("/cart", cartHandler.CreateCart)
	router.Use(app.AuthMiddleware.AuthMiddleware()).GET("/cart", cartHandler.GetCart)
	router.Use(app.AuthMiddleware.AuthMiddleware()).GET("/cart/:id", cartHandler.GetCartByID)
	router.Use(app.AuthMiddleware.AuthMiddleware()).PUT("/cart", cartHandler.UpdateCart)
	router.Use(app.AuthMiddleware.AuthMiddleware()).DELETE("/cart", cartHandler.DeleteCart)

	router.Use(app.AuthMiddleware.AuthMiddleware()).POST("/checkout", checkoutHandler.CreateCheckout)
	router.Use(app.AuthMiddleware.AuthMiddleware()).GET("/checkout", checkoutHandler.GetCurrentCheckout)

	router.Use(app.AuthMiddleware.AuthMiddleware()).POST("/payment", paymentHistoryHandler.CreatePayment)
	router.Use(app.AuthMiddleware.AuthMiddleware()).POST("/payment/cancel", paymentHistoryHandler.CancelPayment)
	router.Use(app.AuthMiddleware.AuthMiddleware()).GET("/payment/history", paymentHistoryHandler.GetPaymentHistory)

	return router
}
