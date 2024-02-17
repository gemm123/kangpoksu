package handler

import (
	"github.com/gin-contrib/sessions"
	"kopoksu/internal/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *dashboardHandler) GetAllProductAdultDiaper(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	diapers, err := h.productService.GetAllProductsAdultDiaper()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-product-adult-diaper.html", gin.H{
		"data":   diapers,
		"status": status,
	})
}

func (h *dashboardHandler) CreateProductAdultDiaper(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	ctx.HTML(http.StatusOK, "dashboard-product-adult-diaper-create.html", gin.H{
		"status": status,
	})
}

func (h *dashboardHandler) PostCreateProductAdultDiaper(ctx *gin.Context) {
	var adultDiaper model.Product

	if err := ctx.ShouldBind(&adultDiaper); err != nil {
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

	adultDiaper.Image = "/static/images/uploads/" + nameFile

	if err := h.productService.SaveProductAdultDiaper(adultDiaper); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/products/adult-diapers")
}

func (h *dashboardHandler) DeleteProductAdultDiaper(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	if err := h.productService.DeleteProduct(id); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/products/adult-diapers")
}

func (h *dashboardHandler) EditProductAdultDiaper(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	product, err := h.productService.EditProduct(id)
	if err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-product-adult-diaper-edit.html", gin.H{
		"data":   product,
		"status": status,
	})
}

func (h *dashboardHandler) UpdateProductAdultDiaper(ctx *gin.Context) {
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

	ctx.Redirect(http.StatusFound, "/dashboard/products/adult-diapers")
}
