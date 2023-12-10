package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *homeHandler) GetAllAdultDiapers(ctx *gin.Context) {
	adultDiapers, err := h.productService.GetAllProductsAdultDiaper()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "adult-diaper.html", gin.H{
		"adultDiapers": adultDiapers,
	})
}

func (h *homeHandler) GetAdultDiaperById(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	adultDiaper, err := h.productService.GetProductById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "adult-diaper-detail.html", gin.H{
		"adultDiaper": adultDiaper,
	})
}
