package cafe

import _entities "potentivio-app/entities"

type CafeUseCaseInterface interface {
	GetCafeById(id int) (_entities.GetCafe, int, error)
	PostCafe(cafe _entities.Cafe) error
	GetAllCafe() ([]_entities.GetAllCafe, error)
	DeleteCafe(id int) error
	UpdateCafe(updateCafe _entities.Cafe, idToken int) (uint, error)
}
