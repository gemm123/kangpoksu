package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kopoksu/helper"
	"kopoksu/internal/model"
	"log"
	"net/http"
	"time"
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

func (h *homeHandler) PostPickupOnlineOrder(ctx *gin.Context) {
	var pickupOnlineOrder model.PickupOnlineOrder
	if err := ctx.ShouldBind(&pickupOnlineOrder); err != nil {
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
	pickupOnlineOrder.Total = totalOrder
	totalFormatted := helper.FormatRupiah(float64(pickupOnlineOrder.Total))

	pickupDate, _ := time.Parse("2006-01-02T15:04", pickupOnlineOrder.PickupDateStr)
	pickupOnlineOrder.PickupDate = pickupDate

	if err := h.pickupOnlineOrderService.SavePickupOnlineOrder(pickupOnlineOrder, cart); err != nil {
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
