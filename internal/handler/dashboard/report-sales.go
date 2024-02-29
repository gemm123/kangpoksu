package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kopoksu/internal/model"
	"log"
	"net/http"
	"strconv"
)

func (h *dashboardHandler) FormReportSalesFormulaMilk(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	ctx.HTML(http.StatusOK, "dashboard-form-sales-formula-milk.html", gin.H{
		"status": status,
	})
}

func (h *dashboardHandler) PostFormReportSalesFormulaMilk(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")
	var productSales []model.ProductSales
	var err error

	monthString := ctx.PostForm("month")
	yearString := ctx.PostForm("year")

	month, _ := strconv.Atoi(monthString)
	year, _ := strconv.Atoi(yearString)

	dayString := ctx.PostForm("day")
	if dayString == "" {
		productSales, err = h.productService.ReportSalesFormulaMilkByMonthYear(month, year)
		if err != nil {
			log.Println("error: " + err.Error())
			return
		}
	} else {
		day, _ := strconv.Atoi(dayString)
		productSales, err = h.productService.ReportSalesFormulaMilkByDate(day, month, year)
		if err != nil {
			log.Println("error: " + err.Error())
			return
		}
	}

	ctx.HTML(http.StatusOK, "dashboard-sales-formula-milk.html", gin.H{
		"status":       status,
		"productSales": productSales,
	})
}

func (h *dashboardHandler) FormReportSalesBabyDiaper(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	ctx.HTML(http.StatusOK, "dashboard-form-sales-baby-diaper.html", gin.H{
		"status": status,
	})
}

func (h *dashboardHandler) PostFormReportSalesBabyDiaper(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")
	var productSales []model.ProductSales
	var err error

	monthString := ctx.PostForm("month")
	yearString := ctx.PostForm("year")

	month, _ := strconv.Atoi(monthString)
	year, _ := strconv.Atoi(yearString)

	dayString := ctx.PostForm("day")
	if dayString == "" {
		productSales, err = h.productService.ReportSalesBabyDiaperByMonthYear(month, year)
		if err != nil {
			log.Println("error: " + err.Error())
			return
		}
	} else {
		day, _ := strconv.Atoi(dayString)
		productSales, err = h.productService.ReportSalesBabyDiaperByDate(day, month, year)
		if err != nil {
			log.Println("error: " + err.Error())
			return
		}
	}

	ctx.HTML(http.StatusOK, "dashboard-sales-baby-diaper.html", gin.H{
		"status":       status,
		"productSales": productSales,
	})
}

func (h *dashboardHandler) FormReportSalesAdultDiaper(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	ctx.HTML(http.StatusOK, "dashboard-form-sales-adult-diaper.html", gin.H{
		"status": status,
	})
}

func (h *dashboardHandler) PostFormReportSalesAdultDiaper(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")
	var productSales []model.ProductSales
	var err error

	monthString := ctx.PostForm("month")
	yearString := ctx.PostForm("year")

	month, _ := strconv.Atoi(monthString)
	year, _ := strconv.Atoi(yearString)

	dayString := ctx.PostForm("day")
	if dayString == "" {
		productSales, err = h.productService.ReportSalesAdultDiaperByMonthYear(month, year)
		if err != nil {
			log.Println("error: " + err.Error())
			return
		}
	} else {
		day, _ := strconv.Atoi(dayString)
		productSales, err = h.productService.ReportSalesAdultDiaperByDate(day, month, year)
		if err != nil {
			log.Println("error: " + err.Error())
			return
		}
	}

	ctx.HTML(http.StatusOK, "dashboard-sales-adult-diaper.html", gin.H{
		"status":       status,
		"productSales": productSales,
	})
}
