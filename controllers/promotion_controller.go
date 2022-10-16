package controllers

import (
	"PromotionsAPI/dto"
	"PromotionsAPI/helpers"
	"PromotionsAPI/initializers"
	"PromotionsAPI/models"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPromotions(context *gin.Context) {

	var promotions models.PromotionList
	result := initializers.DB.Find(&promotions)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, helpers.RESPONSE_NOT_FOUND)
		return
	}

	promotionsListRes := dto.CreatePromotionResponseList(promotions)
	context.JSON(http.StatusOK, promotionsListRes)
}

func GetPromotionById(context *gin.Context) {

	id := context.Param("id")
	var promotion models.Promotion
	result := initializers.DB.First(&promotion, "promotion_id = ?", id)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, helpers.RESPONSE_NOT_FOUND)
		return
	}

	promotionRes := dto.CreatePromotionResponse(promotion)
	context.JSON(http.StatusOK, promotionRes)
}
