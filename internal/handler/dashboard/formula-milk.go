package handler

import (
	"github.com/gin-contrib/sessions"
	"kopoksu/internal/model"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *dashboardHandler) GetAllProductFormulaMilk(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	formulaMilks, err := h.productService.GetAllProductsFormulaMilk()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-product-formula-milk.html", gin.H{
		"data":   formulaMilks,
		"status": status,
	})
}

func (h *dashboardHandler) CreateProductFormulaMilk(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	ctx.HTML(http.StatusOK, "dashboard-product-formula-mlik-create.html", gin.H{
		"status": status,
	})
}

func (h *dashboardHandler) PostCreateProductFormulaMilk(ctx *gin.Context) {
	var formulaMilk model.Product

	if err := ctx.ShouldBind(&formulaMilk); err != nil {
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

	formulaMilk.Image = "/static/images/uploads/" + nameFile

	if err := h.productService.SaveProductFormulaMilk(formulaMilk); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/products/formula-milks")
}

func (h *dashboardHandler) DeleteProductFormulaMilk(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	if err := h.productService.DeleteProduct(id); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/products/formula-milks")
}

func (h *dashboardHandler) EditProductFormulaMilk(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	product, err := h.productService.EditProduct(id)
	if err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-product-formula-milk-edit.html", gin.H{
		"data":   product,
		"status": status,
	})
}

func (h *dashboardHandler) UpdateProductFormulaMilk(ctx *gin.Context) {
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	var newProduct model.Product

	if err := ctx.ShouldBind(&newProduct); err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	file, _ := ctx.FormFile("image")

	if file != nil {
		extension := filepath.Ext(file.Filename)
		nameFile := uuid.NewString() + extension
		destination := "web/static/images/uploads/" + nameFile

		if err := ctx.SaveUploadedFile(file, destination); err != nil {
			log.Println("An error occurred: ", err.Error())
			return
		}

		newProduct.Image = "/static/images/uploads/" + nameFile
	}

	if err := h.productService.UpdateProduct(newProduct, id); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/products/formula-milks")
}
