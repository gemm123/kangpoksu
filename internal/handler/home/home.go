package handler

import (
	"kopoksu/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type homeHandler struct {
	productService      service.ProductService
	cartService         service.CartService
	offlineOrderService service.OfflineOrderService
	onlineOrderService  service.OnlineOrderService
}

func NewHomeHandler(
	productService service.ProductService,
	cartService service.CartService,
	offlineOrderService service.OfflineOrderService,
	onlineOrderService service.OnlineOrderService,
) *homeHandler {
	return &homeHandler{
		productService:      productService,
		cartService:         cartService,
		offlineOrderService: offlineOrderService,
		onlineOrderService:  onlineOrderService,
	}
}

func (h *homeHandler) Home(ctx *gin.Context) {
	formulaMilks, err := h.productService.GetAllProductsFormulaMilkLimit(5)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	babyDiapers, err := h.productService.GetAllProductsBabyDiaperLimit(5)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	adultDiapers, err := h.productService.GetAllProductsAdultDiaperLimit(5)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"formulaMilks": formulaMilks,
		"babyDiapers":  babyDiapers,
		"adultDiapers": adultDiapers,
	})
}
