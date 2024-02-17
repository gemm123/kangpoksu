package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (h *dashboardHandler) GetAllOnlineOrder(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	onlineOrders, err := h.onlineOrderService.GetAllOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-online-order.html", gin.H{
		"data":   onlineOrders,
		"status": status,
	})
}

func (h *dashboardHandler) EditOnlineOrder(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	editOnlineOrderResponse, err := h.onlineOrderService.EditOnlineOrder(id)
	if err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-online-order-edit.html", gin.H{
		"data":   editOnlineOrderResponse,
		"status": status,
	})
}

func (h *dashboardHandler) UpdateOnlineOrder(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	status := ctx.PostForm("status")

	if err := h.onlineOrderService.UpdateStatusOnlineOrder(id, status); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/orders/online")
}

func (h *dashboardHandler) DeleteOnlineOrder(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	if err := h.onlineOrderService.DeleteOnlineOrder(id); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/orders/online")
}
