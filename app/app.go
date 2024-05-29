package app

import (
	"golang-crud-2024/config"
	service "golang-crud-2024/core/service"
	product_category "golang-crud-2024/repository/product-category"
	"golang-crud-2024/repository/user"
)

type App struct {
	AuthService            service.AuthService
	ProductCategoryService service.ProductCategoryService

	Cfg config.Config
}

func NewApp(dep *Dependency) *App {
	userRepo := user.NewUserRepository(dep.Database)
	productCategoryRepo := product_category.NewProductCategoryRepository(dep.Database)

	authService := service.NewAuthService(userRepo, dep.Cfg)
	productCategoryService := service.NewProductCategoryService(productCategoryRepo)

	return &App{
		AuthService:            authService,
		ProductCategoryService: productCategoryService,
		Cfg:                    dep.Cfg,
	}
}
