package handler

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"kopoksu/internal/model"
	"log"
	"net/http"
	"strconv"
)

func GetSessionCart(session sessions.Session, cart []model.Cart) ([]model.Cart, error) {
	cartJson := session.Get("Cart")
	if cartJson != nil {
		if err := json.Unmarshal(cartJson.([]byte), &cart); err != nil {
			return cart, err
		}
	}

	return cart, nil
}

func (h *homeHandler) SaveCartProduct(ctx *gin.Context) {
	var cart []model.Cart
	session := sessions.Default(ctx)

	cart, err := GetSessionCart(session, cart)
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
	var totalOrder int

	cart, err := GetSessionCart(session, cart)
	if err != nil {
		log.Println("error: " + err.Error())
		return
	}

	for i := 0; i < len(cart); i++ {
		product, err := h.productService.GetProductById(cart[i].Id)
		if err != nil {
			log.Println("error: " + err.Error())
			return
		}
		cart[i].Name = product.Name
		cart[i].Image = product.Image
		cart[i].Total = product.Price * cart[i].Amount
		totalOrder = totalOrder + cart[i].Total
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

	cart, err := GetSessionCart(session, cart)
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
