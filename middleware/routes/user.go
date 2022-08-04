package routes

import (
	"chatnews-api/middleware/controller"

	"github.com/labstack/echo/v4"
)

func GetUserApiRoutes(e *echo.Echo, userController *controller.UserController) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("/login", userController.AuthenticateUser)
		v1.GET("/users", userController.GetAllUser)
		v1.POST("/signup", userController.SaveUser)
		v1.GET("/users/:id", userController.GetUser)
		v1.PUT("/users/:id", userController.UpdateUser)
		v1.PUT("/update_password/:id", userController.UpdateUserPassword)
		v1.DELETE("/users/:id", userController.DeleteUser)
		v1.GET("/get_user_by_token", userController.GetUserByToken)
		v1.PUT("/users/edit_id_emp/:id", userController.UpdateUserEmpId)
		v1.GET("/userhod/:id", userController.GetUserHod)
		v1.GET("/userhod/", userController.GetUserHodAll)
		v1.PUT("/edit_section/:id", userController.UpdateUserSection)
		v1.PUT("/edit_position/:id", userController.UpdateUserPosition)
		v1.PUT("/edit_superior/:id", userController.UpdateUserSuperior)
		v1.GET("/usersupervision/:id", userController.GetUserSupervision)
		v1.GET("/usersupervision", userController.GetUserSupervisionAll)
		v1.GET("/username/:id", userController.GetUserName)

	}
}
