package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *homeHandler) GetAllFormulaMilks(ctx *gin.Context) {
	formulaMilks, err := h.productService.GetAllProductsFormulaMilk()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "formula-milk.html", gin.H{
		"formulaMilks": formulaMilks,
	})
}

func (h *homeHandler) GetFormulaMilksById(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	formulaMilk, err := h.productService.GetProductById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "formula-milk-detail.html", gin.H{
		"formulaMilk": formulaMilk,
	})
}
