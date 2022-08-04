package http

import (
	"chatnews-api/api/delivery/controller"
	api "chatnews-api/api/usecase"
	"chatnews-api/middleware/config"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func EmployeeRouter(e *echo.Echo, employeeUseCase api.EmployeeUseCaseI) {

	employeeCtrl := controller.NewEmployeeController(e, employeeUseCase)

	r := e.Group("/api/v1")
	r.Use(middleware.JWT([]byte(config.JWTSecret)))

	r.GET("/employee/list", employeeCtrl.GetData)
	r.POST("/employee/add", employeeCtrl.InsertData)
	r.POST("/employee/edit", employeeCtrl.UpdateData)
	r.POST("/employee/delete", employeeCtrl.DeleteData)

}
