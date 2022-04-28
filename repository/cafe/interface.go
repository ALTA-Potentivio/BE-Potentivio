package cafe

import _entities "potentivio-app/entities"

type CafeRepositoryInterface interface {
	GetCafeById(id int) (_entities.Cafe, int, error)
	PostCafe(cafe _entities.Cafe) error
	GetAllCafe(filters map[string]string) ([]_entities.Cafe, error)
	DeleteCafe(id int) error
	UpdateCafe(updateCafe _entities.Cafe) error
	GetCafeByIdForHire(id uint) (_entities.Cafe, error)
}
