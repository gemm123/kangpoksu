package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kopoksu/helper"
	"kopoksu/internal/model"
	"log"
	"net/http"
)

func (h *homeHandler) FormOrder(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var cart []model.Cart

	cart, err := helper.GetSessionCart(session, cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	totalWeight, err := h.cartService.GetAccumulationTotalWeight(cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "order.html", gin.H{
		"totalWeight": totalWeight,
	})
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

	totalOrder = helper.RandomNumberOrder(totalOrder)
	offlineOrder.Total = totalOrder
	totalFormatted := helper.FormatRupiah(float64(offlineOrder.Total))

	if err := h.offlineOrderService.SaveOfflineOrder(offlineOrder, cart); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	session.Clear()
	session.Save()

	ctx.HTML(http.StatusOK, "payment.html", gin.H{
		"total":          totalOrder,
		"totalFormatted": totalFormatted,
	})
}

func (h *homeHandler) PostOnlineOrder(ctx *gin.Context) {
	var onlineOrder model.OnlineOrder
	if err := ctx.ShouldBind(&onlineOrder); err != nil {
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

	totalOrder = helper.RandomNumberOrder(totalOrder)
	onlineOrder.Total = totalOrder + onlineOrder.Cost
	totalFormatted := helper.FormatRupiah(float64(onlineOrder.Total))

	if err := h.onlineOrderService.SaveOnlineOrder(onlineOrder, cart); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	session.Clear()
	session.Save()

	ctx.HTML(http.StatusOK, "payment.html", gin.H{
		"total":          onlineOrder.Total,
		"totalFormatted": totalFormatted,
	})
}
