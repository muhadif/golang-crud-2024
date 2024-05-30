package app

import (
	"golang-crud-2024/config"
	service "golang-crud-2024/core/service"
	"golang-crud-2024/repository/product"
	product_category "golang-crud-2024/repository/product-category"
	"golang-crud-2024/repository/user"
)

type App struct {
	AuthService            service.AuthService
	ProductCategoryService service.ProductCategoryService
	ProductService         service.ProductService

	Cfg config.Config
}

func NewApp(dep *Dependency) *App {
	userRepo := user.NewUserRepository(dep.Database)
	productCategoryRepo := product_category.NewProductCategoryRepository(dep.Database)
	productRepo := product.NewRepository(dep.Database)

	authService := service.NewAuthService(userRepo, dep.Cfg)
	productCategoryService := service.NewProductCategoryService(productCategoryRepo)
	productService := service.NewProductService(productRepo)

	return &App{
		AuthService:            authService,
		ProductCategoryService: productCategoryService,
		ProductService:         productService,
		Cfg:                    dep.Cfg,
	}
}
