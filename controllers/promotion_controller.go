package controllers

import (
	"PromotionsAPI/dto"
	"PromotionsAPI/helpers"
	"PromotionsAPI/initializers"
	"PromotionsAPI/models"

	"github.com/gin-gonic/gin"

	// "github.com/google/uuid"
	"net/http"
	// "time"
)

// const DATE_TIME_FORMAT string = "2006-01-02 15:04:05"

// func CreatePromotions(c *gin.Context) {

// 	uuid := uuid.New()

// 	promotion := models.Promotion{
// 		UUID:           uuid,
// 		Price:          1.23,
// 		ExpirationDate: time.Now().Format(DATE_TIME_FORMAT),
// 	}

// 	result := initializers.DB.Create(&promotion)

// 	if result.Error != nil {
// 		c.Status(http.StatusNotFound)
// 		return
// 	}

// 	c.JSON(http.StatusOK, promotion)
// }

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
