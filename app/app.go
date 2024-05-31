package app

import (
	service "golang-crud-2024/core/service"
	"golang-crud-2024/pkg/auth"
	"golang-crud-2024/repository/cart"
	"golang-crud-2024/repository/checkout"
	payment_history "golang-crud-2024/repository/payment-history"
	"golang-crud-2024/repository/product"
	product_category "golang-crud-2024/repository/product-category"
	"golang-crud-2024/repository/user"
)

type App struct {
	AuthService            service.AuthService
	ProductCategoryService service.ProductCategoryService
	ProductService         service.ProductService
	CartService            service.CartService
	CheckoutService        service.CheckoutService
	PaymentHistoryService  service.PaymentHistoryService
	PaymentCallbackService service.PaymentCallbackService

	AuthMiddleware auth.Middleware
}

func NewApp(dep *Dependency) *App {
	userRepo := user.NewUserRepository(dep.Database)
	productCategoryRepo := product_category.NewProductCategoryRepository(dep.Database)
	productRepo := product.NewRepository(dep.Database)
	cartRepo := cart.NewRepository(dep.Database)
	checkoutRepo := checkout.NewRepository(dep.Database)
	paymentHistoryRepo := payment_history.NewRepository(dep.Database)

	authService := service.NewAuthService(userRepo, dep.Cfg)
	productCategoryService := service.NewProductCategoryService(productCategoryRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo, productRepo)
	checkoutService := service.NewCheckoutService(checkoutRepo, cartRepo)
	paymentHistoryService := service.NewPaymentHistoryService(paymentHistoryRepo, checkoutService, productRepo)
	paymentCallbackService := service.NewPaymentCallbackService(paymentHistoryRepo)

	authMiddleware := auth.NewMiddleware(dep.Cfg)

	return &App{
		AuthService:            authService,
		ProductCategoryService: productCategoryService,
		ProductService:         productService,
		CartService:            cartService,
		CheckoutService:        checkoutService,
		PaymentHistoryService:  paymentHistoryService,
		PaymentCallbackService: paymentCallbackService,
		AuthMiddleware:         authMiddleware,
	}
}
