package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kopoksu/helper"
	"log"
	"net/http"
)

func (h *dashboardHandler) GetRecapitulation(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	ProfitFormulaMilkPickupOnlineOrder, err := h.recapService.ProfitRecapFormulaMilkPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ProfitBabyDiaperPickupOnlineOrder, err := h.recapService.ProfitRecapBabyDiaperPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ProfitAdultDiaperPickupOnlineOrder, err := h.recapService.ProfitRecapAdultDiaperPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ProfitFormulaMilkOnlineOrder, err := h.recapService.ProfitRecapFormulaMilkOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ProfitBabyDiaperOnlineOrder, err := h.recapService.ProfitRecapBabyDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ProfitAdultDiaperOnlineOrder, err := h.recapService.ProfitRecapAdultDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	recapSalesFormulaMilkByMonthOnlineOrder, err := h.recapService.RecapSalesFormulaMilkByMonthOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	recapSalesBabyDiaperByMonthOnlineOrder, err := h.recapService.RecapSalesBabyDiaperByMonthOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	recapSalesAdultDiaperByMonthOnlineOrder, err := h.recapService.RecapSalesAdultDiaperByMonthOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	recapSalesFormulaMilkByMonthPickupOnlineOrder, err := h.recapService.RecapSalesFormulaMilkByMonthPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	recapSalesBabyDiaperByMonthPickupOnlineOrder, err := h.recapService.RecapSalesBabyDiaperByMonthPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	recapSalesAdultDiaperByMonthPickupOnlineOrder, err := h.recapService.RecapSalesAdultDiaperByMonthPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ProfitFormulaMilk := helper.FormatRupiah(float64(ProfitFormulaMilkPickupOnlineOrder + ProfitFormulaMilkOnlineOrder))
	ProfitBabyDiaper := helper.FormatRupiah(float64(ProfitBabyDiaperPickupOnlineOrder + ProfitBabyDiaperOnlineOrder))
	ProfitAdultDiaper := helper.FormatRupiah(float64(ProfitAdultDiaperPickupOnlineOrder + ProfitAdultDiaperOnlineOrder))

	ctx.HTML(http.StatusOK, "dashboard-recapitulation.html", gin.H{
		"status":            status,
		"ProfitFormulaMilk": ProfitFormulaMilk,
		"ProfitBabyDiaper":  ProfitBabyDiaper,
		"ProfitAdultDiaper": ProfitAdultDiaper,
		"recapSalesFormulaMilkByMonthOnlineOrder":       recapSalesFormulaMilkByMonthOnlineOrder,
		"recapSalesBabyDiaperByMonthOnlineOrder":        recapSalesBabyDiaperByMonthOnlineOrder,
		"recapSalesAdultDiaperByMonthOnlineOrder":       recapSalesAdultDiaperByMonthOnlineOrder,
		"recapSalesFormulaMilkByMonthPickupOnlineOrder": recapSalesFormulaMilkByMonthPickupOnlineOrder,
		"recapSalesBabyDiaperByMonthPickupOnlineOrder":  recapSalesBabyDiaperByMonthPickupOnlineOrder,
		"recapSalesAdultDiaperByMonthPickupOnlineOrder": recapSalesAdultDiaperByMonthPickupOnlineOrder,
	})
}
