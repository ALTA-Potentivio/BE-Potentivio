package hire

import (
	"potentivio-app/entities"
)

type HireUseCaseInterface interface {
	CreateHire(hire entities.Hire) error
	GetHireByIdArtist(IdArtist int) ([]entities.Hire, error)
	GetHireByIdCafe(IdCafe int) ([]entities.Hire, error)
	AcceptHire(hire entities.Hire) error
	CancelHireByCafe(hire entities.Hire) error
	Rejecthire(hire entities.Hire) error
	CancelHireByArtis(hire entities.Hire) error
	Rating(hire entities.Hire) error
	CallBack(hire entities.Hire) error
	Done(hire entities.Hire) error
}
