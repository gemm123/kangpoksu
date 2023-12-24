package handler

import (
	"github.com/gin-gonic/gin"
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
