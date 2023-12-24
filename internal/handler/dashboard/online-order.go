package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (h *dashboardHandler) GetAllOnlineOrder(ctx *gin.Context) {
	onlineOrders, err := h.onlineOrderService.GetAllOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-online-order.html", gin.H{
		"data": onlineOrders,
	})
}

func (h *dashboardHandler) EditOnlineOrder(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	editOnlineOrderResponse, err := h.onlineOrderService.EditOnlineOrder(id)
	if err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-online-order-edit.html", gin.H{
		"data": editOnlineOrderResponse,
	})
}
