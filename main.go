package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/tugasecho/controllers"
)

func main() {
	//Buat instance baru dari Echo dengan fungsi New()
	e := echo.New()

	//Menggunakan Use() untuk menambahkan middleware yang akan menyimpan koneksi database ke dalam konteks Echo
	//sehingga dapat diakses oleh handler
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", controllers.Connect())
			return next(c)
		}
	})

	//Mendefinisikan route HTTP dan handler
	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.InsertUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	//Menjalankan server Echo dengan Start() dan dijalankan pada port 6969
	//Menangani error dengan Fatal() serta menampilkan pesan log dan menghentikan app jika terjadi error
	e.Logger.Fatal(e.Start(":6969"))
}
