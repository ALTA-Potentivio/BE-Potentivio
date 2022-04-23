package cafe

import _entities "potentivio-app/entities"

type CafeUseCaseInterface interface {
	GetCafeById(id int) (_entities.Cafe, int, error)
}
