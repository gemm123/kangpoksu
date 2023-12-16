package helper

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"kopoksu/internal/model"
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
