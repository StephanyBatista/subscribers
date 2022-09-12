package users

import (
	"log"
	"net/http"
	"subscribers/domain"
	"subscribers/web"
	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserRepository Repository
}

func (h *Handler) Post(c *gin.Context) {
	var body UserRequest
	c.BindJSON(&body)
	errs := domain.Validate(body)
	if errs != nil {
		log.Println(errs)
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	user, err := NewUser(body.Name, body.Email, body.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, web.NewErrorReponse(err.Error()))
		return
	}

	userSaved, _ := h.UserRepository.GetByEmail(body.Email)
	if userSaved.Id != "" {
		log.Println(body.Email + " already exist")
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("Email already saved"))
		return
	}

	err = h.UserRepository.Create(user)
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
	var body UserChangePasswordRequest
	c.BindJSON(&body)
	token := c.GetHeader("Authorization")
	claim, _ := auth.GetClaimFromToken(token)
	userSaved, _ := h.UserRepository.GetByEmail(claim.Email)

	err := userSaved.ChangePassword(body.OldPassword, body.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorReponse(err.Error()))
		return
	}
	h.UserRepository.Save(userSaved)
	c.JSON(http.StatusOK, http.StatusOK)
}

func (h *Handler) Token(c *gin.Context) {
	var body LoginRequest
	c.BindJSON(&body)
	errs := domain.Validate(body)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
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
