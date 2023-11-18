package handler

import (
	"net/http"
	"template/internal/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ApiGetMessage(ctx *gin.Context) {
	message := h.service.GetMessage()
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (h *handler) WebGetMessage(ctx *gin.Context) {
	_ = h.service.GetMessage()
	ctx.HTML(http.StatusOK, "home.html", gin.H{})
}
