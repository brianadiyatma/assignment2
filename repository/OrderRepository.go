package repository

import (
	"second-assignment/model"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) Create(newPerson model.Order) (model.Order, error) {
	tx := or.db.Create(&newPerson)
	return newPerson, tx.Error
}

func (or *orderRepository) GetAll() ([]model.Order, error) {
	var persons = []model.Order{}
	tx := or.db.Preload("Items").Find(&persons)
	return persons, tx.Error
}

func (or *orderRepository) Delete(uuid string) error {
	tx := or.db.Delete(&model.Order{}, "order_id= ?", uuid)
	return tx.Error
}

func (or *orderRepository) Update(val *model.Order) ( error) {
		tx := or.db.Save(val)
		return  tx.Error
}

