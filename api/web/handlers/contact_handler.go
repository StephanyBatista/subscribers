package handlers

import (
	"net/http"
	"subscribers/domain"
	"subscribers/domain/contacts"
	"subscribers/domain/users"
	"subscribers/infra/database"
	"subscribers/web"

	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	ContactRepository database.IRepository[contacts.Contact]
	UserRepository    database.IRepository[users.User]
}

func (h *ContactHandler) Post(c *gin.Context) {
	var body ContactRequest
	c.BindJSON(&body)
	errs := domain.Validate(body)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity := contacts.NewContact(body.Name, body.Email, claim.UserId)
	ok := h.ContactRepository.Create(entity)
	if !ok {
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": entity.ID})
}

func (h *ContactHandler) GetAll(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entities := h.ContactRepository.List(contacts.Contact{UserId: claim.UserId})
	c.JSON(http.StatusOK, entities)
}

func (h *ContactHandler) GetById(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))
	id := c.Param("id")

	entity := h.ContactRepository.GetBy(contacts.Contact{Entity: domain.Entity{ID: id}, UserId: claim.UserId})
	if entity == nil {
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
	}
	c.JSON(http.StatusOK, entity)
}
