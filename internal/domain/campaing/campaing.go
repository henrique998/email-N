package campaing

import (
	"time"

	internalerrors "github.com/henrique1501/email-N/internal/internalErrors"
	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaing struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreatedAt time.Time `validate:"required"`
}

func NewCampaing(name, content string, emails []string) (*Campaing, error) {
	contacts := make([]Contact, len(emails))

	for i, email := range emails {
		contacts[i].Email = email
	}

	campaing := &Campaing{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}

	err := internalerrors.ValidateStruct(campaing)
	if err != nil {
		return nil, err
	}

	return campaing, nil
}
