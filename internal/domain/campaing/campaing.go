package campaing

import (
	"time"

	internalerrors "github.com/henrique998/email-N/internal/internalErrors"
	"github.com/rs/xid"
)

const (
	Pending = "Pending"
	Started = "Started"
	Done    = "Done"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaing struct {
	ID        string    `validate:"required" json:"id"`
	Name      string    `validate:"min=5,max=24" json:"name"`
	Content   string    `validate:"min=5,max=1024" json:"content"`
	Contacts  []Contact `validate:"min=1,dive" json:"contacts"`
	CreatedAt time.Time `validate:"required" json:"createdAt"`
	Status    string
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
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaing)
	if err != nil {
		return nil, err
	}

	return campaing, nil
}
