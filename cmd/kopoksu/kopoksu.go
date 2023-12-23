package main

import (
	"kopoksu/config"
	dashboardHandler "kopoksu/internal/handler/dashboard"
	homeHandler "kopoksu/internal/handler/home"
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
	orderRepository := repository.NewOrderRepository(db)

	//Service
	adminService := service.NewAdminService()
	productService := service.NewProductService(productRepository)
	cartService := service.NewCartService(productRepository)
	orderService := service.NewOrderService(orderRepository, productRepository)

	//Handler
	dashboardHandler := dashboardHandler.NewDashboardHandler(adminService, productService, orderService)
	homeHandler := homeHandler.NewHomeHandler(productService, cartService, orderService)

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/**/*")

	router.Static("/static", "./web/static")

	authStore := cookie.NewStore([]byte("auth-secret"))
	authStore.Options(sessions.Options{
		MaxAge:   7 * 24 * 60 * 60,
		Path:     "/dashboard",
		HttpOnly: true,
	})

	cartStore := cookie.NewStore([]byte("cart-secret"))
	cartStore.Options(sessions.Options{
		MaxAge:   3 * 24 * 60 * 60,
		Path:     "/",
		HttpOnly: true,
	})

	dashboard := router.Group("/dashboard")
	dashboard.Use(sessions.Sessions("auth-session", authStore))
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

	dashboard.GET("/products/adult-diapers", dashboardHandler.GetAllProductAdultDiaper)
	dashboard.GET("/products/adult-diapers/create", dashboardHandler.CreateProductAdultDiaper)
	dashboard.POST("/products/adult-diapers/create", dashboardHandler.PostCreateProductAdultDiaper)
	dashboard.POST("/products/adult-diapers/delete/:id", dashboardHandler.DeleteProductAdultDiaper)
	dashboard.GET("/products/adult-diapers/edit/:id", dashboardHandler.EditProductAdultDiaper)
	dashboard.POST("/products/adult-diapers/edit/:id", dashboardHandler.UpdateProductAdultDiaper)

	dashboard.GET("/orders/offline", dashboardHandler.GetAllOfflineOrder)
	dashboard.GET("/orders/offline/edit/:id", dashboardHandler.EditOfflineOrder)
	dashboard.POST("/orders/offline/edit/:id", dashboardHandler.UpdateOfflineOrder)

	router.Use(sessions.Sessions("cart-session", cartStore))
	router.GET("/", homeHandler.Home)
	router.GET("/formula-milks", homeHandler.GetAllFormulaMilks)
	router.GET("/formula-milks/:id", homeHandler.GetFormulaMilksById)

	router.GET("/baby-diapers", homeHandler.GetAllBabyDiapers)
	router.GET("/baby-diapers/:id", homeHandler.GetBabyDiaperById)

	router.GET("/adult-diapers", homeHandler.GetAllAdultDiapers)
	router.GET("/adult-diapers/:id", homeHandler.GetAdultDiaperById)

	router.POST("/add-cart", homeHandler.SaveCartProduct)
	router.GET("/cart", homeHandler.GetCartProduct)
	router.POST("/cart-delete", homeHandler.DeleteProductAtCart)
	router.GET("/order", homeHandler.FormOrder)
	router.POST("/order/offline", homeHandler.PostOfflineOrder)
	router.POST("/order/online", homeHandler.PostOnlineOrder)

	router.Run()
}
