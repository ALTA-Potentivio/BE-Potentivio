package hire

import (
	"potentivio-app/entities"
)

type HireUseCaseInterface interface {
	CreateHire(hire entities.Hire) error
	GetHireByIdArtist(IdArtist int) ([]entities.Hire, error)
	GetHireByIdCafe(IdCafe int) ([]entities.Hire, error)
	AcceptHire(hire entities.Hire) error
}
