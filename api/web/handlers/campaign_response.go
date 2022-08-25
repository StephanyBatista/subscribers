package handlers

type CampaignResponse struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Status            string `json:"status"`
	From              string `json:"from"`
	Subject           string `json:"subject"`
	Body              string `json:"body"`
	AttachmentURL     string `json:"attachmentURL"`
	BaseOfSubscribers int    `json:"baseofSubscribers"`
	TotalSent         int    `json:"totalSent"`
	TotalRead         int    `json:"totalRead"`
}
