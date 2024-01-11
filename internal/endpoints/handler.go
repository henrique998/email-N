package endpoints

import "github.com/henrique1501/email-N/internal/domain/campaing"

type Handler struct {
	CampaignService campaing.Service
}
