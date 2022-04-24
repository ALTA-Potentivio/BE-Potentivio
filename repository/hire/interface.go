package hire

import (
	"potentivio-app/entities"
)

type HireRepositoryInterface interface {
	CheckHire(hire entities.Hire) error
	AcceptHire(hire entities.Hire) error
	CreateHire(hire entities.Hire) error
	GetHireByIdArtis(IdArtist int) ([]entities.Hire, error)
	GetHireByIdCafe(IdCafe int) ([]entities.Hire, error)
}
