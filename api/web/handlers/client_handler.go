package handlers

import (
	"net/http"
	"subscribers/domain"
	"subscribers/domain/clients"
	"subscribers/domain/users"
	"subscribers/infra/database"
	"subscribers/web"

	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	ClientRepository database.IRepository[clients.Client]
	UserRepository   database.IRepository[users.User]
}

func (h *ClientHandler) Post(c *gin.Context) {
	var body ClientRequest
	c.BindJSON(&body)
	errs := domain.Validate(body)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity := clients.NewClient(body.Name, body.Email, claim.UserId)
	ok := h.ClientRepository.Create(entity)
	if !ok {
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": entity.ID})
}

func (h *ClientHandler) GetAll(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entities := h.ClientRepository.List(clients.Client{UserId: claim.UserId})
	c.JSON(http.StatusOK, entities)
}

func (h *ClientHandler) GetById(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))
	id := c.Param("id")

	entity := h.ClientRepository.GetBy(clients.Client{Entity: domain.Entity{ID: id}, UserId: claim.UserId})
	if entity == nil {
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
	}
	c.JSON(http.StatusOK, entity)
}
