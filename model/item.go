package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	Id          string `gorm:"column:item_id;primaryKey" json:"id"`
	ItemCode    string `gorm:"column:item_code" json:"item_code"`
	Description string `gorm:"column:description" json:"description"`
	Quantity    int    `gorm:"column:quantity" json:"quantity"`
	OrderID     string `gorm:"column:order_id;foreignKey" json:"order_id"`
}

func (o *Item) BeforeCreate(tx *gorm.DB) error {

	o.Id = uuid.NewString()
	
	return nil
}