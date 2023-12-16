package handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kopoksu/helper"
	"kopoksu/internal/model"
	"log"
	"net/http"
)

func (h *homeHandler) FormOrder(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "order.html", gin.H{})
}

func (h *homeHandler) PostOfflineOrder(ctx *gin.Context) {
	var offlineOrder model.OfflineOrder
	if err := ctx.ShouldBind(&offlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	session := sessions.Default(ctx)
	var cart []model.Cart

	cart, err := helper.GetSessionCart(session, cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	totalOrder, err := h.cartService.GetAccumulationTotalCart(cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	fmt.Println(offlineOrder)
	fmt.Println(cart)
	fmt.Println(totalOrder)
}

func (h *homeHandler) PostOnlineOrder(ctx *gin.Context) {
	var onlineOrder model.OnlineOrder
	if err := ctx.ShouldBind(&onlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	fmt.Println(onlineOrder)
}
