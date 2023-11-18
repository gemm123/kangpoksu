package main

import (
	"log"
	"template/config"
	"template/internal/handler"
	"template/internal/repository"
	"template/internal/service"

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

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", handler.ApiGetMessage)
	router.GET("/web", handler.WebGetMessage)
	router.Run()
}
