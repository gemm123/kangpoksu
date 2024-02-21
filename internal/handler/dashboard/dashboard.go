package handler

import (
	"kopoksu/internal/model"
	"kopoksu/internal/service"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type dashboardHandler struct {
	adminService        service.AdminService
	productService      service.ProductService
	offlineOrderService service.OfflineOrderService
	onlineOrderService  service.OnlineOrderService
	recapService        service.RecapService
}

func NewDashboardHandler(
	adminService service.AdminService,
	productService service.ProductService,
	offlineOrderService service.OfflineOrderService,
	onlineOrderService service.OnlineOrderService,
	recapService service.RecapService,
) *dashboardHandler {
	return &dashboardHandler{
		adminService:        adminService,
		productService:      productService,
		offlineOrderService: offlineOrderService,
		onlineOrderService:  onlineOrderService,
		recapService:        recapService,
	}
}

func SaveUploadFile(file *multipart.FileHeader, ctx *gin.Context) (string, error) {
	extension := filepath.Ext(file.Filename)
	nameFile := uuid.NewString() + extension
	destination := "web/static/images/uploads/" + nameFile

	if err := ctx.SaveUploadedFile(file, destination); err != nil {
		log.Println("An error occurred: ", err.Error())
		return "", err
	}

	return nameFile, nil
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

	if strings.Contains(admin.Email, "master") {
		session.Set("user", "master")
	} else if strings.Contains(admin.Email, "admin") {
		session.Set("user", "admin")
	}

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
	session := sessions.Default(ctx)
	status := session.Get("user")

	countOfflineOrderConfirmationPayment, err := h.offlineOrderService.CountOfflineOrderByStatus("Menunggu konfirmasi pembayaran")
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	countOfflineOrderTake, err := h.offlineOrderService.CountOfflineOrderByStatus("Menunggu pengambilan")
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	countOfflineOrderDone, err := h.offlineOrderService.CountOfflineOrderByStatus("Selesai")
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	countOnlineOrderConfirmationPayment, err := h.onlineOrderService.CountOnlineOrderByStatus("Menunggu konfirmasi pembayaran")
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	countOnlineOrderDelivery, err := h.onlineOrderService.CountOnlineOrderByStatus("Pengiriman")
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	countOnlineOrderDone, err := h.onlineOrderService.CountOnlineOrderByStatus("Selesai")
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	totalCountConfirmationPayment := countOfflineOrderConfirmationPayment + countOnlineOrderConfirmationPayment
	totalCountDelivery := countOnlineOrderDelivery
	totalCountTake := countOfflineOrderTake
	totalCountDone := countOfflineOrderDone + countOnlineOrderDone

	ctx.HTML(http.StatusOK, "dashboard-home.html", gin.H{
		"confirmationPayment": totalCountConfirmationPayment,
		"countDone":           totalCountDone,
		"countDelivery":       totalCountDelivery,
		"countTake":           totalCountTake,
		"status":              status,
	})
}
