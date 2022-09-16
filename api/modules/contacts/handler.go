package contacts

import (
	"fmt"
	"net/http"
	"subscribers/commun/web"
	"subscribers/commun/web/auth"
	"subscribers/modules/users"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ContactRepository Repository
	UserRepository    users.Repository
}

func (h *Handler) Post(c *gin.Context) {
	var body CreateNewContact
	c.BindJSON(&body)

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity, errs := NewContact(body.Name, body.Email, claim.UserId)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}
	err := h.ContactRepository.Create(entity)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": entity.Id})
}

func (h *Handler) GetAll(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entities, _ := h.ContactRepository.ListBy(claim.UserId)
	c.JSON(http.StatusOK, entities)
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")

	entity, _ := h.ContactRepository.GetBy(id)
	if entity.Id == "" {
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	c.JSON(http.StatusOK, entity)
}

func (h *Handler) Cancel(c *gin.Context) {
	id := c.Param("id")

	entity, _ := h.ContactRepository.GetBy(id)
	if entity.Id == "" {
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
	}
	entity.Cancel()
	//TODO: validate err and create test
	h.ContactRepository.Save(entity)
	c.JSON(http.StatusOK, http.StatusOK)
}
