package handlers

import (
	"log"
	"net/http"
	"subscribers/domain"
	"subscribers/domain/clients"
	"subscribers/domain/users"
	"subscribers/web"

	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientHandler struct {
	Db *gorm.DB
}

func (h *ClientHandler) Post(c *gin.Context) {
	var body ClientRequest
	c.BindJSON(&body)
	errs := domain.Validate(body)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	userId := c.Param("userId")

	var userFound users.User
	result := h.Db.Where(users.User{Entity: domain.Entity{ID: userId}}).Find(&userFound)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("User not found"))
		return
	}

	entity := clients.NewClient(body.Name, body.Email, userId)
	result = h.Db.Create(&entity)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": entity.ID})
}

func (h *ClientHandler) GetAll(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	var entities []clients.Client
	result := h.Db.Where(clients.Client{UserId: claim.UserId}).Find(&entities)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusOK, entities)
}

func (h *ClientHandler) GetById(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))
	id := c.Param("id")

	var entity clients.Client
	result :=
		h.Db.Where(clients.Client{Entity: domain.Entity{ID: id}, UserId: claim.UserId}).Find(&entity)
	if result.RowsAffected == 0 {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}

	c.JSON(http.StatusOK, entity)
}
