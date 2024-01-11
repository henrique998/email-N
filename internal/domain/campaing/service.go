package campaing

import (
	"github.com/henrique1501/email-N/internal/contracts"
	internalerrors "github.com/henrique1501/email-N/internal/internalErrors"
)

type Service struct {
	Repo Repository
}

func (s *Service) Create(newCampaing contracts.NewCampaingDTO) (string, error) {
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
