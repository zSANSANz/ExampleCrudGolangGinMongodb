package http

import (
	"chatnews-api/api/delivery/controller"
	api "chatnews-api/api/usecase"
	"chatnews-api/middleware/config"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func PerformanceRouter(e *echo.Echo, performanceUseCase api.PerformanceUseCaseI) {

	performanceCtrl := controller.NewPerformanceController(e, performanceUseCase)

	r := e.Group("/api/v1")
	r.Use(middleware.JWT([]byte(config.JWTSecret)))

	r.GET("/performance/list", performanceCtrl.GetData)
	r.GET("/performance/list_like_employee", performanceCtrl.GetDataUCUser)
	r.GET("/performance/list_like_devision", performanceCtrl.GetDataUCDevision)
	r.GET("/performance/list_by_updated_at", performanceCtrl.GetDataUCOrderByUpdatedAt)
	r.POST("/performance/add", performanceCtrl.InsertData)
	r.POST("/performance/edit", performanceCtrl.UpdateData)
	r.POST("/performance/edit_status", performanceCtrl.UpdateStatus)
	r.POST("/performance/delete", performanceCtrl.DeleteData)

}
