package users

import (
	"errors"
	"log"
	"net/http"
	"subscribers/web"
	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserRepository Repository
}

func (h *Handler) Post(c *gin.Context) {
	var body CreateNewUser
	c.BindJSON(&body)

	user, errs := NewUser(body.Name, body.Email, body.Password)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	userSaved, _ := h.UserRepository.GetByEmail(body.Email)
	if userSaved.Id != "" {
		log.Println(body.Email + " already exist")
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("Email already saved"))
		return
	}

	err := h.UserRepository.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": user.Id})
}

func (h *Handler) GetInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claim, ok := auth.GetClaimFromToken(token)
	if !ok {
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": claim.UserName, "email": claim.Email})
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var body ChangePassword
	c.BindJSON(&body)

	token := c.GetHeader("Authorization")
	claim, _ := auth.GetClaimFromToken(token)
	userSaved, _ := h.UserRepository.GetByEmail(claim.Email)

	errs := userSaved.ChangePassword(body.OldPassword, body.NewPassword)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}
	h.UserRepository.Save(userSaved)
	c.JSON(http.StatusOK, http.StatusOK)
}

func (h *Handler) Token(c *gin.Context) {
	var body Login
	c.BindJSON(&body)
	var errs []error
	if body.Email == "" {
		errs = append(errs, errors.New("'Email' is required"))
	}
	if body.Password == "" {
		errs = append(errs, errors.New("'Password' is required"))
	}
	if len(errs) > 0 {
		c.JSON(http.StatusForbidden, web.NewErrorsReponse(errs))
		return
	}

	userSaved, _ := h.UserRepository.GetByEmail(body.Email)
	if userSaved.Id == "" || !userSaved.CheckPassword(body.Password) {
		log.Println(body.Email + " not found")
		c.JSON(http.StatusForbidden, web.NewErrorReponse("User not found  or password invalid"))
		return
	}

	token, expiresAt, err := auth.GenerateJWT(userSaved.Id, userSaved.Email, userSaved.Name)
	if err != nil {
		log.Println("Error to generate JWT")
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "expiresAt": expiresAt})
}
