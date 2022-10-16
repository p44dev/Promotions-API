package routers

import (
	"PromotionsAPI/controllers"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func ConfigureRoutes() {

	router := gin.Default()
	metricRouter := gin.Default()
	monitor := ginmetrics.GetMonitor()

	// Configure Metrics
	monitor.UseWithoutExposingEndpoint(router)
	monitor.Expose(metricRouter)
	monitor.SetMetricPath("/metrics")
	monitor.SetSlowTime(10)
	monitor.Use(metricRouter)

	router.GET("/promotions", controllers.GetPromotions)
	router.GET("/promotions/:id", controllers.GetPromotionById)

	go func() {
		metricRouter.Run(":9091")
	}()

	router.Run()
}
