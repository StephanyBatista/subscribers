package clients

type CreationRequest struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	CampaignId string `json:"campaignId" validate:"required"`
}
