package contacts

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscribers/domain"
	"subscribers/modules/users"
	"subscribers/web"
	"subscribers/web/auth"
)

type Handler struct {
	ContactRepository Repository
	UserRepository    users.Repository
}

func (h *Handler) Post(c *gin.Context) {
	var body ContactRequest
	c.BindJSON(&body)
	errs := domain.Validate(body)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity := NewContact(body.Name, body.Email, claim.UserId)
	err := h.ContactRepository.Create(entity)
	if err != nil {
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
	h.ContactRepository.Save(entity)
	c.JSON(http.StatusOK, http.StatusOK)
}
