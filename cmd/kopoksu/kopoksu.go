package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"template/config"
	"template/internal/handler"
	"template/internal/service"
	"template/middleware"
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

	adminService := service.NewAdminService()
	dashboardHandler := handler.NewDashboardHandler(adminService)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

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

	router.Run()
}
