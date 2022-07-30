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

type UserHandler struct {
	Db *gorm.DB
}

type UserCreationRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *UserHandler) Post(c *gin.Context) {
	var body UserCreationRequest
	c.BindJSON(&body)
	if !web.Validate(body, c) {
		return
	}

	var userSaved user.User
	h.Db.Where(user.User{Email: body.Email}).First(&userSaved)
	if !userSaved.IDIsNull() {
		log.Println(body.Email + " already exist")
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("Email already saved"))
		return
	}

	user, err := user.NewUser(body.Name, body.Email, body.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	result := h.Db.Create(&user)
	if result.Error != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": user.ID})
}

func (h *UserHandler) GetInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claim, ok := auth.GetClaimFromToken(token)
	if !ok {
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Name": claim.Name, "Email": claim.Email})
}
