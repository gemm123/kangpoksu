package handler

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"kopoksu/helper"
	"kopoksu/internal/model"
	"log"
	"net/http"
	"strconv"
)

func (h *homeHandler) SaveCartProduct(ctx *gin.Context) {
	var cart []model.Cart
	session := sessions.Default(ctx)

	cart, err := helper.GetSessionCart(session, cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	idString := ctx.PostForm("id")
	id := uuid.MustParse(idString)

	amountString := ctx.PostForm("amount")
	amount, _ := strconv.Atoi(amountString)

	product := model.Cart{
		Id:     id,
		Amount: amount,
	}
	cart = append(cart, product)

	cartJson, err := json.Marshal(cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	session.Set("Cart", cartJson)
	if err := session.Save(); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/cart")
}

func (h *homeHandler) GetCartProduct(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var cart []model.Cart

	cart, err := helper.GetSessionCart(session, cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	totalOrder, err := h.cartService.GetAccumulationTotalCart(cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "cart.html", gin.H{
		"cart":       cart,
		"totalOrder": totalOrder,
	})
}

func (h *homeHandler) DeleteProductAtCart(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var cart []model.Cart
	var updatedCart []model.Cart

	cart, err := helper.GetSessionCart(session, cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	idString := ctx.PostForm("id")
	id := uuid.MustParse(idString)

	for _, c := range cart {
		if c.Id != id {
			updatedCart = append(updatedCart, c)
		}
	}

	updatedCartJson, err := json.Marshal(updatedCart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	session.Set("Cart", updatedCartJson)
	if err := session.Save(); err != nil {
		log.Println("error: " + err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/cart")
}
