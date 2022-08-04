package routes

import (
	"chatnews-api/middleware/controller"

	"github.com/labstack/echo/v4"
)

func GetPerformanceApiRoutes(e *echo.Echo, performanceController *controller.PerformanceController) {
	v1 := e.Group("/api/v1")
	{
		v1.GET("/performance_all", performanceController.GetAllPerformance)
		v1.GET("/performancehod/:id", performanceController.GetPerformanceHod)
		v1.GET("/performancesupervision/:id", performanceController.GetPerformanceSupervision)
		v1.GET("/performanceuser/:id", performanceController.GetPerformanceUser)
		v1.GET("/performanceuserid/:id", performanceController.GetPerformanceUserId)
		v1.GET("/performanceid/:id", performanceController.GetPerformanceId)
	}
}
