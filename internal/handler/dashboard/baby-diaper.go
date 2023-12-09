package handler

import (
	"kopoksu/internal/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *dashboardHandler) GetAllProductBabyDiaper(ctx *gin.Context) {
	diapers, err := h.productService.GetAllProductsBabyDiaper()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-product-baby-diaper.html", gin.H{
		"data": diapers,
	})
}

func (h *dashboardHandler) CreateProductBabyDiaper(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "dashboard-product-baby-diaper-create.html", gin.H{})
}

func (h *dashboardHandler) PostCreateProductBabyDiaper(ctx *gin.Context) {
	var babyDiaper model.Product

	if err := ctx.ShouldBind(&babyDiaper); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	nameFile, err := SaveUploadFile(file, ctx)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	babyDiaper.Image = "/static/images/uploads/" + nameFile

	if err := h.productService.SaveProductBabyDiaper(babyDiaper); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/products/baby-diapers")
}

func (h *dashboardHandler) DeleteProductBabyDiaper(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	if err := h.productService.DeleteProduct(id); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/products/baby-diapers")
}

func (h *dashboardHandler) EditProductBabyDiaper(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	product, err := h.productService.EditProduct(id)
	if err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-product-baby-diaper-edit.html", gin.H{
		"data": product,
	})
}

func (h *dashboardHandler) UpdateProductBabyDiaper(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	var newProduct model.Product

	if err := ctx.ShouldBind(&newProduct); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	file, _ := ctx.FormFile("image")

	if file != nil {
		nameFile, err := SaveUploadFile(file, ctx)
		if err != nil {
			log.Println("error: " + err.Error())
			return
		}

		newProduct.Image = "/static/images/uploads/" + nameFile
	}

	if err := h.productService.UpdateProduct(newProduct, id); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/products/baby-diapers")
}
