package handlers

import (
	"log"
	"net/http"
	"subscribers/domain/users"
	"subscribers/web"
	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	Db *gorm.DB
}

func (h *UserHandler) Post(c *gin.Context) {
	var body users.CreationRequest
	c.BindJSON(&body)

	user, errs := users.NewUser(body)
	if errs != nil {
		log.Println(errs)
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	var userSaved users.User
	h.Db.Where(users.User{Email: body.Email}).First(&userSaved)
	if !userSaved.IDIsNull() {
		log.Println(body.Email + " already exist")
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("Email already saved"))
		return
	}

	result := h.Db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
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
	c.JSON(http.StatusOK, gin.H{"Name": claim.UserName, "Email": claim.Email})
}
