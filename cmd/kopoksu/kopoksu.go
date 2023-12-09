package main

import (
	"kopoksu/config"
	handler "kopoksu/internal/handler/dashboard"
	"kopoksu/internal/repository"
	"kopoksu/internal/service"
	"kopoksu/middleware"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env")
	}

	db, err := config.NewDbPool()
	if err != nil {
		log.Fatal(err)
	}
	defer config.CloseDB(db)

	//Repository
	productRepository := repository.NewProductRepository(db)

	//Service
	adminService := service.NewAdminService()
	productService := service.NewProductService(productRepository)

	//Handler
	dashboardHandler := handler.NewDashboardHandler(adminService, productService)

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/**/*")

	router.Static("/static", "./web/static")

	store := cookie.NewStore([]byte("kopoksu"))
	store.Options(sessions.Options{
		MaxAge:   7 * 24 * 60 * 60,
		Path:     "/",
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("mysession", store))

	dashboard := router.Group("/dashboard")
	dashboard.GET("/login", dashboardHandler.Login)
	dashboard.POST("/login", dashboardHandler.PostLogin)
	dashboard.POST("/logout", dashboardHandler.Logout)

	dashboard.Use(middleware.AuthAdmin())
	dashboard.GET("/home", dashboardHandler.Home)
	dashboard.GET("/products/formula-milks", dashboardHandler.GetAllProductFormulaMilk)
	dashboard.GET("/products/formula-milks/create", dashboardHandler.CreateProductFormulaMilk)
	dashboard.POST("/products/formula-milks/create", dashboardHandler.PostCreateProductFormulaMilk)
	dashboard.GET("/products/formula-milks/edit/:id", dashboardHandler.EditProductFormulaMilk)
	dashboard.POST("/products/formula-milks/edit/:id", dashboardHandler.UpdateProductFormulaMilk)
	dashboard.POST("/products/formula-milks/delete/:id", dashboardHandler.DeleteProductFormulaMilk)

	dashboard.GET("/products/baby-diapers", dashboardHandler.GetAllProductBabyDiaper)
	dashboard.GET("/products/baby-diapers/create", dashboardHandler.CreateProductBabyDiaper)
	dashboard.POST("/products/baby-diapers/create", dashboardHandler.PostCreateProductBabyDiaper)
	dashboard.POST("/products/baby-diapers/delete/:id", dashboardHandler.DeleteProductBabyDiaper)
	dashboard.GET("/products/baby-diapers/edit/:id", dashboardHandler.EditProductBabyDiaper)
	dashboard.POST("/products/baby-diapers/edit/:id", dashboardHandler.UpdateProductBabyDiaper)

	router.Run()
}
