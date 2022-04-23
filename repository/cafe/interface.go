package cafe

import _entities "potentivio-app/entities"

type CafeRepositoryInterface interface {
	GetCafeById(id int) (_entities.Cafe, int, error)
}
