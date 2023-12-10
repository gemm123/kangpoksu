package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *homeHandler) GetAllBabyDiapers(ctx *gin.Context) {
	babyDiapers, err := h.productService.GetAllProductsBabyDiaper()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "baby-diaper.html", gin.H{
		"babyDiapers": babyDiapers,
	})
}

func (h *homeHandler) GetBabyDiaperById(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	babyDiaper, err := h.productService.GetProductById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "baby-diaper-detail.html", gin.H{
		"babyDiaper": babyDiaper,
	})
}
