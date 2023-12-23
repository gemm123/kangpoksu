package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (h *dashboardHandler) GetAllOfflineOrder(ctx *gin.Context) {
	offlineOrders, err := h.offlineOrderService.GetAllOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-offline-order.html", gin.H{
		"data": offlineOrders,
	})
}

func (h *dashboardHandler) EditOfflineOrder(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	editOfflineOrderResponse, err := h.offlineOrderService.EditOfflineOrder(id)
	if err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-offline-order-edit.html", gin.H{
		"data": editOfflineOrderResponse,
	})
}

func (h *dashboardHandler) UpdateOfflineOrder(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	status := ctx.PostForm("status")

	if err := h.offlineOrderService.UpdateStatusOfflineOrder(id, status); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/orders/offline")
}

func (h *dashboardHandler) DeleteOfflineOrder(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	if err := h.offlineOrderService.DeleteOfflineOrder(id); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/orders/offline")
}
