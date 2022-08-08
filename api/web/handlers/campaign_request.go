package handlers

type CampaignRequest struct {
	Name string `json:"name" validate:"required"`
	From string `json:"from" validate:"required"`
	Body string `json:"body" validate:"required"`
}
