package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	Id          string `gorm:"column:item_id;primaryKey"`
	ItemCode    string `gorm:"column:item_code"`
	Description string `gorm:"column:description"`
	Quantity    int    `gorm:"column:quantity"`
	OrderID     string `gorm:"column:order_id;foreignKey"`
}

func (o *Item) BeforeCreate(tx *gorm.DB) error {

	o.Id = uuid.NewString()
	
	return nil
}