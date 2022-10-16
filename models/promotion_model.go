package models

import (
	"github.com/google/uuid"
)

type Promotion struct {
	PromotionId    uint      `gorm:"column:promotion_id;autoIncrement;primary_key"`
	UUID           uuid.UUID `gorm:"colomn:id;type:uuid"`
	Price          float64   `gorm:"colomn:price"`
	ExpirationDate string    `gorm:"colomn:expiration_date"`
}

type PromotionList []*Promotion

func (t *Promotion) TableName() string {
	return "promotions"
}
