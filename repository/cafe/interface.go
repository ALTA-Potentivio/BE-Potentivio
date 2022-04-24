package cafe

import _entities "potentivio-app/entities"

type CafeRepositoryInterface interface {
	GetCafeById(id int) (_entities.Cafe, int, error)
	PostCafe(cafe _entities.Cafe) error
	GetAllCafe() ([]_entities.Cafe, error)
	DeleteCafe(id int) error
	UpdateCafe(updateCafe _entities.Cafe, idToken int) (uint, error)
}
