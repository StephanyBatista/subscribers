package campaigns

type CreationRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Active      bool   `json:"active" validate:"required"`
}
