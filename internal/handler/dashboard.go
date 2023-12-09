package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"template/internal/model"
	"template/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type dashboardHandler struct {
	adminService   service.AdminService
	productService service.ProductService
}

func NewDashboardHandler(adminService service.AdminService, productService service.ProductService) *dashboardHandler {
	return &dashboardHandler{
		adminService:   adminService,
		productService: productService,
	}
}

func (h *dashboardHandler) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "dashboard-login.html", gin.H{})
}

func (h *dashboardHandler) PostLogin(ctx *gin.Context) {
	var admin model.AdminLogin
	session := sessions.Default(ctx)

	if err := ctx.Bind(&admin); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ok := h.adminService.CheckCredentials(admin.Email, admin.Password)
	if !ok {
		log.Println("wrong credentials")
		ctx.HTML(http.StatusOK, "dashboard-login.html", gin.H{})
		return
	}

	log.Println("success login")
	session.Set("login", true)
	session.Save()

	ctx.Redirect(http.StatusFound, "/dashboard/home")
}

func (h *dashboardHandler) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Clear()
	session.Save()

	ctx.Redirect(http.StatusFound, "/dashboard/login")
}

func (h *dashboardHandler) Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "dashboard-home.html", gin.H{})
}

func (h *dashboardHandler) GetAllProductFormulaMilk(ctx *gin.Context) {
	formulaMilks, err := h.productService.GetAllProductsFormulaMilk()
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-product-formula-milk.html", gin.H{
		"data": formulaMilks,
	})
}

func (h *dashboardHandler) CreateProductFormulaMilk(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "dashboard-product-formula-mlik-create.html", gin.H{})
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

	extension := filepath.Ext(file.Filename)
	nameFile := uuid.NewString() + extension
	destination := "web/static/images/uploads/" + nameFile

	if err := ctx.SaveUploadedFile(file, destination); err != nil {
		log.Println("An error occurred: ", err.Error())
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
	idString := ctx.Param("id")
	id := uuid.MustParse(idString)

	product, err := h.productService.EditProduct(id)
	if err != nil {
		log.Println("An error occurred: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-product-formula-milk-edit.html", gin.H{
		"data": product,
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
