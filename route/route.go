package route

import (
	"echo-book/controller"

	"github.com/labstack/echo"
)


func New() *echo.Echo  {
e := echo.New()
	// Route / to handler function
e.GET("/users", controller.GetUsersController)
e.GET("/users/:id", controller.GetUserController)
e.POST("/users", controller.CreateUserController)
e.DELETE("/users/:id", controller.DeleteUserController)
e.PUT("/users/:id", controller.UpdateUserController)

return e

}

