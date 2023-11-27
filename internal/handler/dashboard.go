package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"template/internal/model"
	"template/internal/service"
)

type dashboardHandler struct {
	adminService service.AdminService
}

func NewDashboardHandler(adminService service.AdminService) *dashboardHandler {
	return &dashboardHandler{
		adminService: adminService,
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
