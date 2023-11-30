package handler

import (
	"log"
	"net/http"
	"template/internal/model"
	"template/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	return
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
		return
	}

	log.Println("success login")
	session.Set("login", true)
	session.Save()

	ctx.Redirect(http.StatusFound, "/dashboard/home")

	return
}

func (h *dashboardHandler) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Clear()
	session.Save()

	ctx.Redirect(http.StatusFound, "/dashboard/login")

	return
}

func (h *dashboardHandler) Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "dashboard-home.html", gin.H{})
	return
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

	return
}
