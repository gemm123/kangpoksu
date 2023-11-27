package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		statusLogin := session.Get("login")
		if statusLogin != true {
			log.Println("login first!")
			ctx.Redirect(http.StatusFound, "/dashboard/login")
			return
		}

		ctx.Next()
	}
}
