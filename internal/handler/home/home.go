package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kopoksu/internal/service"
	"log"
	"net/http"
)

type homeHandler struct {
	productService           service.ProductService
	cartService              service.CartService
	pickupOnlineOrderService service.PickupOnlineOrderService
	onlineOrderService       service.OnlineOrderService
}

func NewHomeHandler(
	productService service.ProductService,
	cartService service.CartService,
	pickupOnlineOrderService service.PickupOnlineOrderService,
	onlineOrderService service.OnlineOrderService,
) *homeHandler {
	return &homeHandler{
		productService:           productService,
		cartService:              cartService,
		pickupOnlineOrderService: pickupOnlineOrderService,
		onlineOrderService:       onlineOrderService,
	}
}

func (h *homeHandler) Home(ctx *gin.Context) {
	formulaMilks, err := h.productService.GetAllProductsFormulaMilkLimit(4)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	babyDiapers, err := h.productService.GetAllProductsBabyDiaperLimit(4)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	adultDiapers, err := h.productService.GetAllProductsAdultDiaperLimit(4)
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

type SearchResult struct {
	Results string `json:"results"`
}

func (h *homeHandler) Search(ctx *gin.Context) {
	search := ctx.Query("term")

	searchResult, err := h.productService.SearchProductsByName(search)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	fmt.Println(searchResult)

	ctx.JSON(http.StatusOK, searchResult)
}
