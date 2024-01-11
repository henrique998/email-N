package database

import "github.com/henrique1501/email-N/internal/domain/campaing"

type CampaignRepository struct {
	campaigns []campaing.Campaing
}

func (c *CampaignRepository) Save(campaing *campaing.Campaing) error {
	c.campaigns = append(c.campaigns, *campaing)

	return nil
}

func (c *CampaignRepository) Get() ([]campaing.Campaing, error) {
	return c.campaigns, nil
}
