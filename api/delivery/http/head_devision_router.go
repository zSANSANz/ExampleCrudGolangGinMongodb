package http

import (
	"chatnews-api/api/delivery/controller"
	api "chatnews-api/api/usecase"
	"chatnews-api/middleware/config"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func HeadDevisionRouter(e *echo.Echo, headDevisionUseCase api.HeadDevisionUseCaseI) {

	headDevisionCtrl := controller.NewHeadDevisionController(e, headDevisionUseCase)

	r := e.Group("/api/v1")
	r.Use(middleware.JWT([]byte(config.JWTSecret)))

	r.GET("/head_devision/list", headDevisionCtrl.GetData)
	r.POST("/head_devision/add", headDevisionCtrl.InsertData)
	r.POST("/head_devision/edit", headDevisionCtrl.UpdateData)
	r.POST("/head_devision/delete", headDevisionCtrl.DeleteData)

}
