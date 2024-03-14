package repository

import (
	"second-assignment/model"
)


type IOrderRepository interface {
	Create(newPerson model.Order) (model.Order, error)
	GetAll() ([]model.Order, error)
	Delete(uuid string) error
	Update(val *model.Order) error
}