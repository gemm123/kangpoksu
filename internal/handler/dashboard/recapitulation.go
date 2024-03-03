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

	//grossProfitFormulaMilkOfflineOrder, err := h.recapService.GrossProfitRecapFormulaMilkOfflineOrder()
	//if err != nil {
	//	log.Println("error: " + err.Error())
	//	return
	//}
	//
	//grossProfitBabyDiaperOfflineOrder, err := h.recapService.GrossProfitRecapBabyDiaperOfflineOrder()
	//if err != nil {
	//	log.Println("error: " + err.Error())
	//	return
	//}
	//
	//grossProfitAdultDiaperOfflineOrder, err := h.recapService.GrossProfitRecapAdultDiaperOfflineOrder()
	//if err != nil {
	//	log.Println("error: " + err.Error())
	//	return
	//}
	//
	//grossProfitFormulaMilkOnlineOrder, err := h.recapService.GrossProfitRecapFormulaMilkOnlineOrder()
	//if err != nil {
	//	log.Println("error: " + err.Error())
	//	return
	//}
	//
	//grossProfitBabyDiaperOnlineOrder, err := h.recapService.GrossProfitRecapBabyDiaperOnlineOrder()
	//if err != nil {
	//	log.Println("error: " + err.Error())
	//	return
	//}
	//
	//grossProfitAdultDiaperOnlineOrder, err := h.recapService.GrossProfitRecapAdultDiaperOnlineOrder()
	//if err != nil {
	//	log.Println("error: " + err.Error())
	//	return
	//}

	ProfitFormulaMilkOfflineOrder, err := h.recapService.ProfitRecapFormulaMilkOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ProfitBabyDiaperOfflineOrder, err := h.recapService.ProfitRecapBabyDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ProfitAdultDiaperOfflineOrder, err := h.recapService.ProfitRecapAdultDiaperOfflineOrder()
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

	recapSalesFormulaMilkByMonthOfflineOrder, err := h.recapService.RecapSalesFormulaMilkByMonthOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	recapSalesBabyDiaperByMonthOfflineOrder, err := h.recapService.RecapSalesBabyDiaperByMonthOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	recapSalesAdultDiaperByMonthOfflineOrder, err := h.recapService.RecapSalesAdultDiaperByMonthOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	//grossProfitFormulaMilk := helper.FormatRupiah(float64(grossProfitFormulaMilkOfflineOrder + grossProfitFormulaMilkOnlineOrder))
	//grossProfitBabyDiaper := helper.FormatRupiah(float64(grossProfitBabyDiaperOfflineOrder + grossProfitBabyDiaperOnlineOrder))
	//grossProfitAdultDiaper := helper.FormatRupiah(float64(grossProfitAdultDiaperOfflineOrder + grossProfitAdultDiaperOnlineOrder))

	ProfitFormulaMilk := helper.FormatRupiah(float64(ProfitFormulaMilkOfflineOrder + ProfitFormulaMilkOnlineOrder))
	ProfitBabyDiaper := helper.FormatRupiah(float64(ProfitBabyDiaperOfflineOrder + ProfitBabyDiaperOnlineOrder))
	ProfitAdultDiaper := helper.FormatRupiah(float64(ProfitAdultDiaperOfflineOrder + ProfitAdultDiaperOnlineOrder))

	ctx.HTML(http.StatusOK, "dashboard-recapitulation.html", gin.H{
		"status": status,
		//"grossProfitFormulaMilk":                   grossProfitFormulaMilk,
		//"grossProfitBabyDiaper":                    grossProfitBabyDiaper,
		//"grossProfitAdultDiaper":                   grossProfitAdultDiaper,
		"ProfitFormulaMilk":                        ProfitFormulaMilk,
		"ProfitBabyDiaper":                         ProfitBabyDiaper,
		"ProfitAdultDiaper":                        ProfitAdultDiaper,
		"recapSalesFormulaMilkByMonthOnlineOrder":  recapSalesFormulaMilkByMonthOnlineOrder,
		"recapSalesBabyDiaperByMonthOnlineOrder":   recapSalesBabyDiaperByMonthOnlineOrder,
		"recapSalesAdultDiaperByMonthOnlineOrder":  recapSalesAdultDiaperByMonthOnlineOrder,
		"recapSalesFormulaMilkByMonthOfflineOrder": recapSalesFormulaMilkByMonthOfflineOrder,
		"recapSalesBabyDiaperByMonthOfflineOrder":  recapSalesBabyDiaperByMonthOfflineOrder,
		"recapSalesAdultDiaperByMonthOfflineOrder": recapSalesAdultDiaperByMonthOfflineOrder,
	})
}
