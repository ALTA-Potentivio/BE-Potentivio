package hire

import (
	"potentivio-app/entities"
)

type HireRepositoryInterface interface {
	CheckHire(hire entities.Hire) entities.Hire
	AcceptHire(hire entities.Hire) error
	CreateHire(hire entities.Hire) error
	GetHireByIdArtist(IdArtist int) ([]entities.Hire, error)
	GetHireById(id int) (entities.Hire, error)
	GetHireByIdCafe(IdCafe int) ([]entities.Hire, error)
	UpdateHire(id int, hire entities.Hire) error
	DeleteHire(hire entities.Hire) error
	Rating(rating entities.Rating) error
	CallBack(hire entities.Hire) error
}
