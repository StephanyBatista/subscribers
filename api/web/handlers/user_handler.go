package handlers

import (
	"log"
	"net/http"
	"subscribers/domain/user"
	"subscribers/web"

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
	h.Db.Where(user.User{Email: body.Email}).FirstOrInit(&userSaved)
	if userSaved.ID > 0 {
		log.Println(body.Email + " already exist")
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("Email already existe"))
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

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{Db: db}
}
