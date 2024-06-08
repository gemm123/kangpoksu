package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (h *dashboardHandler) GetAllPickupOnlineOrder(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	pickupOnlineOrders, err := h.pickupOnlineOrderService.GetAllPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-pickup-online-order.html", gin.H{
		"data":   pickupOnlineOrders,
		"status": status,
	})
}

func (h *dashboardHandler) EditPickupOnlineOrder(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	editPickupOnlineOrderResponse, err := h.pickupOnlineOrderService.EditPickupOnlineOrder(id)
	if err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-pickup-online-order-edit.html", gin.H{
		"data":   editPickupOnlineOrderResponse,
		"status": status,
	})
}

func (h *dashboardHandler) UpdatePickupOnlineOrder(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	status := ctx.PostForm("status")

	if err := h.pickupOnlineOrderService.UpdateStatusPickupOnlineOrder(id, status); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/orders/pickup-online")
}

func (h *dashboardHandler) DeletePickupOnlineOrder(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	if err := h.pickupOnlineOrderService.DeletePickupOnlineOrder(id); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/orders/pickup-online")
}
