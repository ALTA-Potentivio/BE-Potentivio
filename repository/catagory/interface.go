package catagory

import (
	_entities "potentivio-app/entities"
)

type CatagoryRepositoryInterface interface {
	CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error)
	GetAllCatagory() ([]_entities.Catagory, error)
}
