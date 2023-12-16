package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kopoksu/internal/model"
	"log"
	"net/http"
)

func (h *homeHandler) FormOrder(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "order.html", gin.H{})
}

func (h *homeHandler) PostOfflineOrder(ctx *gin.Context) {
	var offlineOrder model.OfflineOrder
	if err := ctx.ShouldBind(&offlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	fmt.Println(offlineOrder)
}

func (h *homeHandler) PostOnlineOrder(ctx *gin.Context) {
	var onlineOrder model.OnlineOrder
	if err := ctx.ShouldBind(&onlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	fmt.Println(onlineOrder)
}
