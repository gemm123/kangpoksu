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

	grossProfitFormulaMilkOfflineOrder, err := h.recapService.GrossProfitRecapFormulaMilkOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	grossProfitBabyDiaperOfflineOrder, err := h.recapService.GrossProfitRecapBabyDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	grossProfitAdultDiaperOfflineOrder, err := h.recapService.GrossProfitRecapAdultDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	grossProfitFormulaMilkOnlineOrder, err := h.recapService.GrossProfitRecapFormulaMilkOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	grossProfitBabyDiaperOnlineOrder, err := h.recapService.GrossProfitRecapBabyDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	grossProfitAdultDiaperOnlineOrder, err := h.recapService.GrossProfitRecapAdultDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	netProfitFormulaMilkOfflineOrder, err := h.recapService.NetProfitRecapFormulaMilkOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	netProfitBabyDiaperOfflineOrder, err := h.recapService.NetProfitRecapBabyDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	netProfitAdultDiaperOfflineOrder, err := h.recapService.NetProfitRecapAdultDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	netProfitFormulaMilkOnlineOrder, err := h.recapService.NetProfitRecapFormulaMilkOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	netProfitBabyDiaperOnlineOrder, err := h.recapService.NetProfitRecapBabyDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	netProfitAdultDiaperOnlineOrder, err := h.recapService.NetProfitRecapAdultDiaperOnlineOrder()
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

	grossProfitFormulaMilk := helper.FormatRupiah(float64(grossProfitFormulaMilkOfflineOrder + grossProfitFormulaMilkOnlineOrder))
	grossProfitBabyDiaper := helper.FormatRupiah(float64(grossProfitBabyDiaperOfflineOrder + grossProfitBabyDiaperOnlineOrder))
	grossProfitAdultDiaper := helper.FormatRupiah(float64(grossProfitAdultDiaperOfflineOrder + grossProfitAdultDiaperOnlineOrder))

	netProfitFormulaMilk := helper.FormatRupiah(float64(netProfitFormulaMilkOfflineOrder + netProfitFormulaMilkOnlineOrder))
	netProfitBabyDiaper := helper.FormatRupiah(float64(netProfitBabyDiaperOfflineOrder + netProfitBabyDiaperOnlineOrder))
	netProfitAdultDiaper := helper.FormatRupiah(float64(netProfitAdultDiaperOfflineOrder + netProfitAdultDiaperOnlineOrder))

	ctx.HTML(http.StatusOK, "dashboard-recapitulation.html", gin.H{
		"status":                                  status,
		"grossProfitFormulaMilk":                  grossProfitFormulaMilk,
		"grossProfitBabyDiaper":                   grossProfitBabyDiaper,
		"grossProfitAdultDiaper":                  grossProfitAdultDiaper,
		"netProfitFormulaMilk":                    netProfitFormulaMilk,
		"netProfitBabyDiaper":                     netProfitBabyDiaper,
		"netProfitAdultDiaper":                    netProfitAdultDiaper,
		"recapSalesFormulaMilkByMonthOnlineOrder": recapSalesFormulaMilkByMonthOnlineOrder,
		"recapSalesBabyDiaperByMonthOnlineOrder":  recapSalesBabyDiaperByMonthOnlineOrder,
		"recapSalesAdultDiaperByMonthOnlineOrder": recapSalesAdultDiaperByMonthOnlineOrder,
	})
}
