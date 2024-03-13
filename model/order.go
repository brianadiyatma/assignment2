package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	OrderId string `gorm:"column:order_id;primaryKey"`
	OrderedAt time.Time `gorm:"column:ordered_at"`
	CustomerName string `gorm:"column:customer_name"`
	Items []Item 
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {

	o.OrderId = uuid.NewString()
	o.OrderedAt = time.Now()
	return nil
}