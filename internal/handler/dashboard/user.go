package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"kopoksu/internal/model"
	"log"
	"net/http"
)

func (h *dashboardHandler) GetAllUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	users, err := h.userService.GetAllUser()
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-user.html", gin.H{
		"data":   users,
		"status": status,
	})
}

func (h *dashboardHandler) EditUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	id := ctx.Param("id")
	user, err := h.userService.GetUserById(id)
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "dashboard-user-edit.html", gin.H{
		"data":   user,
		"status": status,
	})
}

func (h *dashboardHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.PostForm("user-name")
	email := ctx.PostForm("user-email")
	password := ctx.PostForm("user-password")

	user := model.User{
		Id:       uuid.MustParse(id),
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := h.userService.UpdateUser(user); err != nil {
		log.Println("error: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/management-users")
}

func (h *dashboardHandler) AddUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get("user")

	ctx.HTML(http.StatusOK, "dashboard-user-create.html", gin.H{
		"status": status,
	})
}

func (h *dashboardHandler) PostAddUser(ctx *gin.Context) {
	name := ctx.PostForm("user-name")
	email := ctx.PostForm("user-email")
	password := ctx.PostForm("user-password")

	user := model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := h.userService.CreateUser(user); err != nil {
		log.Println("error: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/management-users")
}

func (h *dashboardHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.userService.DeleteUser(id); err != nil {
		log.Println("error: ", err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard/management-users")
}
