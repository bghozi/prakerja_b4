package routes

import (
	"prakerja_b4/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.InsertUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	return e
}
