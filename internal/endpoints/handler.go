package endpoints

import "github.com/henrique998/email-N/internal/domain/campaing"

type Handler struct {
	CampaignService campaing.Service
}
