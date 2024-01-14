package database

import (
	"github.com/henrique998/email-N/internal/domain/campaing"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Save(campaing *campaing.Campaing) error {
	tx := c.Db.Save(campaing)

	return tx.Error
}

func (c *CampaignRepository) Get() ([]campaing.Campaing, error) {
	var campaigns []campaing.Campaing

	tx := c.Db.Find(&campaigns)

	return campaigns, tx.Error
}

func (c *CampaignRepository) GetById(campaignId string) (*campaing.Campaing, error) {
	var campaign campaing.Campaing

	tx := c.Db.First(&campaign, "id = ?", campaignId)

	return &campaign, tx.Error
}
