package handlers

import (
	"log"
	"net/http"
	"subscribers/domain/user"
	"subscribers/web"
	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TokenHandler struct {
	Db *gorm.DB
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *TokenHandler) Post(c *gin.Context) {
	var body LoginRequest
	c.BindJSON(&body)
	if !web.Validate(body, c) {
		return
	}

	var userSaved user.User
	h.Db.Where(user.User{Email: body.Email}).FirstOrInit(&userSaved)
	if userSaved.ID == 0 || !userSaved.CheckPassword(body.Password) {
		log.Println(body.Email + " not found")
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("User not found"))
		return
	}

	token, expiresAt, err := auth.GenerateJWT(userSaved.Email, userSaved.Name)
	if err != nil {
		log.Println("Error to generate JWT")
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "expiresAt": expiresAt})
}
