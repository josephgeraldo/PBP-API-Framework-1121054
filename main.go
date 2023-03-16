package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/tugasecho/controllers"
)

func main() {
	e := echo.New()

	// Menyimpan koneksi database ke dalam context Echo
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", controllers.Connect())
			return next(c)
		}
	})

	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.InsertUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":6969"))
}
