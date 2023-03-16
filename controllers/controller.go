package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	db := c.Get("db").(*sql.DB)

	query := "SELECT id, name, age, address, country FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": "Something Wrong",})
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Country); err != nil {
			log.Println("Error:", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"Message": "Something Wrong",})
		} else {
			users = append(users, user)
		}
	}
	
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