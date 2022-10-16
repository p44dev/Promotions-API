package dto

import (
	"PromotionsAPI/models"
	"github.com/google/uuid"
)

type PromotionResponse struct {
	UUID           uuid.UUID `json:"id"`
	Price          float64   `json:"price"`
	ExpirationDate string    `json:"expiration_date"`
}

func CreatePromotionResponse(promotion models.Promotion) PromotionResponse {
	return PromotionResponse{
		UUID:           promotion.UUID,
		Price:          promotion.Price,
		ExpirationDate: promotion.ExpirationDate,
	}
}

type PromotionsListResponse []*PromotionResponse

func CreatePromotionResponseList(promotions models.PromotionList) PromotionsListResponse {
	promotionsResp := PromotionsListResponse{}
	for _, p := range promotions {
		promotion := CreatePromotionResponse(*p)
		promotionsResp = append(promotionsResp, &promotion)
	}
	return promotionsResp
}
