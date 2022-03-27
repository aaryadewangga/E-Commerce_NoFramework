package app

import "P2/app/models"

type Model struct {
	Model interface{}
}

func RegistryModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Section{}},
		{Model: models.Category{}},
	}
}
