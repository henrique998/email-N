package campaing

import (
	"errors"
	"testing"

	"github.com/henrique998/email-N/internal/contracts"
	internalerrors "github.com/henrique998/email-N/internal/internalErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(campaing *Campaing) error {
	args := r.Called(campaing)

	return args.Error(0)
}

func (r *repositoryMock) Update(campaing *Campaing) error {
	args := r.Called(campaing)

	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaing, error) {

	return nil, nil
}

func (r *repositoryMock) GetById(campaingId string) (*Campaing, error) {
	args := r.Called(campaingId)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Campaing), nil
}

func (r *repositoryMock) Cancel(campaignId string) error {

	return nil
}

func (r *repositoryMock) Delete(campaign *Campaing) error {

	return nil
}

var (
	newCampaing = contracts.NewCampaingDTO{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"jhondoe@gmail.com", "henrique@gmail.com"},
	}
	service = ServiceImp{}
)

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	repository.On("Save", mock.Anything).Return(nil)
	service.Repo = repository

	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contracts.NewCampaingDTO{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Save_CreateCampaing(t *testing.T) {
	repository := new(repositoryMock)
	repository.On("Save", mock.MatchedBy(func(campaing *Campaing) bool {
		if campaing.Name != newCampaing.Name {
			return false
		}

		if campaing.Content != newCampaing.Content {
			return false
		}

		if len(campaing.Contacts) != len(newCampaing.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repo = repository

	service.Create(newCampaing)

	repository.AssertExpectations(t)
}

func Test_Save_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	repository.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repo = repository

	_, err := service.Create(newCampaing)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_FindById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repository := new(repositoryMock)
	repository.On("GetById", mock.MatchedBy(func(campaignId string) bool {
		return campaignId == campaign.ID
	})).Return(campaign, nil)
	service.Repo = repository

	campaignReturned, _ := service.Repo.GetById(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
}

func Test_FindById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repository := new(repositoryMock)
	repository.On("GetById", mock.Anything).Return(nil, errors.New("Something wrong"))
	service.Repo = repository

	_, err := service.FindById(campaign.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}
