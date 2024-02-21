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

	grossProfitFormulaMilk := helper.FormatRupiah(float64(grossProfitFormulaMilkOfflineOrder + grossProfitFormulaMilkOnlineOrder))
	grossProfitBabyDiaper := helper.FormatRupiah(float64(grossProfitBabyDiaperOfflineOrder + grossProfitBabyDiaperOnlineOrder))
	grossProfitAdultDiaper := helper.FormatRupiah(float64(grossProfitAdultDiaperOfflineOrder + grossProfitAdultDiaperOnlineOrder))

	ctx.HTML(http.StatusOK, "dashboard-recapitulation.html", gin.H{
		"status":                 status,
		"grossProfitFormulaMilk": grossProfitFormulaMilk,
		"grossProfitBabyDiaper":  grossProfitBabyDiaper,
		"grossProfitAdultDiaper": grossProfitAdultDiaper,
	})
}
