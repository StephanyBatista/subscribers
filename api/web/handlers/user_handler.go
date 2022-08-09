package handlers

import (
	"log"
	"net/http"
	"subscribers/domain"
	"subscribers/domain/users"
	"subscribers/infra/database"
	"subscribers/web"
	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepository database.IRepository[users.User]
}

func (h *UserHandler) Post(c *gin.Context) {
	var body UserRequest
	c.BindJSON(&body)
	errs := domain.Validate(body)
	if errs != nil {
		log.Println(errs)
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	user, err := users.NewUser(body.Name, body.Email, body.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, web.NewErrorReponse(err.Error()))
		return
	}

	userSaved := h.UserRepository.GetBy(users.User{Email: body.Email})
	if userSaved != nil {
		log.Println(body.Email + " already exist")
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("Email already saved"))
		return
	}

	ok := h.UserRepository.Create(user)
	if !ok {
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
	c.JSON(http.StatusOK, gin.H{"name": claim.UserName, "email": claim.Email})
}
