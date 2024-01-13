package campaing

type Repository interface {
	Save(campaing *Campaing) error
	Get() ([]Campaing, error)
	GetById(campaignId string) (*Campaing, error)
}
