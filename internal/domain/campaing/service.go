package campaing

import (
	"errors"

	"github.com/henrique998/email-N/internal/contracts"
	internalerrors "github.com/henrique998/email-N/internal/internalErrors"
)

type Service interface {
	Create(newCampaing contracts.NewCampaingDTO) (string, error)
	FindById(campaingId string) (*contracts.CampaignResponseDTO, error)
	Cancel(campaingId string) error
}

type ServiceImp struct {
	Repo Repository
}

func (s *ServiceImp) Create(newCampaing contracts.NewCampaingDTO) (string, error) {
	campaing, err := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repo.Save(campaing)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaing.ID, nil
}

func (s *ServiceImp) FindById(campaingId string) (*contracts.CampaignResponseDTO, error) {
	campaign, err := s.Repo.GetById(campaingId)
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contracts.CampaignResponseDTO{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}

func (s *ServiceImp) Cancel(campaingId string) error {
	campaign, err := s.Repo.GetById(campaingId)
	if err != nil {
		return internalerrors.ErrInternal
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.Cancel()

	err = s.Repo.Save(campaign)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
