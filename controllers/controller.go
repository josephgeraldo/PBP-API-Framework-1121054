package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//echo.Context digunakan untuk mengambil data dari request
func GetAllUsers(c echo.Context) error {
	//Mengambil koneksi database dari konteks Echo menggunakan c.Get dan diubah menjadi tipe data *sql.DB
	db := c.Get("db").(*sql.DB)
	defer db.Close()

	//Membuat query untuk mengambil semua data user
	query := "SELECT id, name, age, address, country FROM users"

	//Mengambil data dari query
	name := c.QueryParam("name")
	ages, ok := c.QueryParams()["age"]

	//Mengecek apakah ada query parameter name
	if name != "" {
  		fmt.Println(name)
  		query += " WHERE name='" + name + "'"
	}

	//Mengecek apakah ada query parameter age
	if ok && len(ages) > 0 && ages[0] != "" {
  		if name != "" {
    		query += " AND"
  		} else {
    		query += " WHERE"
  		}
 		query += " age='" + ages[0] + "'"
	}

	//Mengeksekusi query
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": "Something Wrong",})
	}

	//Membuat slice untuk menampung data user
	var users []User
	for rows.Next() {
		var user User
		//Mengambil data dari query dan memasukkannya ke struct User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Country); err != nil {
			log.Println("Error:", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"Message": "Something Wrong",})
		} else {
			users = append(users, user)
		}
	}
	
	//Mengecek apakah data user ada jika ada maka akan mengembalikan data user
	//jika tidak maka akan mengembalikan pesan error
	if len(users) != 0 {
		var response UsersResponse
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = users
		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusOK, response)
	} else {
		var response ErrorResponse
		response.Status = http.StatusBadRequest
		response.Message = "Error"
		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusBadRequest, response)
	}
}

func InsertUser(c echo.Context) error {
	//Mengambil koneksi database dari konteks Echo menggunakan c.Get dan diubah menjadi tipe data *sql.DB
	db := c.Get("db").(*sql.DB)
	defer db.Close()

	//Mengambil data dari form dan mengecek apakah terjadi kesalahan saat pemrosesan form
	err := c.Request().ParseForm()
	if err != nil {
		log.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": "Something Wrong",})
	}

	//Mengambil data dari form dan ditampung ke variabel
	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	country := c.FormValue("country")

	//Mengeksekusi query untuk menambahkan data user
	ins, errQuery := db.Exec("INSERT INTO users (name, age, address, country) VALUES (?, ?, ?, ?)", name, age, address, country)
	//Mengambil id dari data user yang baru ditambahkan
	id, errQuery := ins.LastInsertId()
	
	//Membuat struct User untuk menampung data user yang baru ditambahkan
	user := User{
		ID: int(id),
		Name: name,
		Age: age,
		Address: address,
		Country: country,
	}
	
	//Mengecek apakah terjadi kesalahan saat menambahkan data user
	var response Response
	if errQuery != nil {
		response.Message = "Error"
		return c.JSONP(http.StatusBadRequest, response.Message, user)
	} else {
		response.Message = "Success"
		return c.JSONP(http.StatusOK, response.Message, user)
	}
}

func UpdateUser(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	defer db.Close()

	//Mengambil data id dari url
	id, _ := strconv.Atoi(c.Param("id"))

	//Mengecek apakah id user ada
	err := c.Request().ParseForm()
	if err != nil {
		log.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": "Something Wrong",})
	}

	//Mengambil data dari form dan ditampung ke variabel
	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	country := c.FormValue("country")

	//Mengeksekusi query untuk mengubah data user
	ins, errQuery := db.Exec("UPDATE users SET name=?, age=?, address=?, country=? WHERE id=?", name, age, address, country, id)
	
	//Mengecek apakah terdapat id user yang diubah
	rowsAffected, _ := ins.RowsAffected()
	if rowsAffected == 0 {
		response := ResponseDelete{
			Status : 400,
			Message: "User not found",
			Id     : id,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	//Mengambil id dari data user yang baru ditambahkan
	_, errQuery = ins.LastInsertId()
	
	//Membuat struct User untuk menampung data user yang baru ditambahkan
	user := User{
		ID: id,
		Name: name,
		Age: age,
		Address: address,
		Country: country,
	}
	
	//Mengecek apakah terjadi kesalahan saat menambahkan data user
	var response Response
	if errQuery != nil {
		response.Message = "Error"
		return c.JSONP(http.StatusBadRequest, response.Message, user)
	} else {
		response.Message = "Success"
		return c.JSONP(http.StatusOK, response.Message, user)
	}
}

func DeleteUser(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	defer db.Close()

	//Mengambil data id dari url
	id, _ := strconv.Atoi(c.Param("id"))

	//Mengeksekusi query untuk menghapus data user
	res, errQuery := db.Exec("DELETE FROM users WHERE id=?", id)
	
	//Mengecek apakah data user ada dengan RowsAffected()
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		response := ResponseDelete{
			Status : 400,
			Message: "User not found",
			Id     : id,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	//Mengatur ulang auto increment setelah data user dihapus
	_, errAutoIncrement := db.Exec(fmt.Sprintf("ALTER TABLE users AUTO_INCREMENT = %d", id))
	if errAutoIncrement != nil {
		response := ResponseDelete{
			Status : 400,
			Message: "Error",
			Id     : id,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	//Mengecek apakah terjadi kesalahan saat menghapus data user
	var response ResponseDelete
	if errQuery != nil {
		response.Status = 400
		response.Message = "Error"
		response.Id = id
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Status = 200
		response.Message = "Success Delete User"
		response.Id = id
		return c.JSON(http.StatusOK, response)
	}
}