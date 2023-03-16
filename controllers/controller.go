package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	defer db.Close()

	query := "SELECT id, name, age, address, country FROM users"

	name := c.QueryParam("name")
	ages, ok := c.QueryParams()["age"]
	if name != "" {
  		fmt.Println(name)
  		query += " WHERE name='" + name + "'"
	}
	if ok && len(ages) > 0 && ages[0] != "" {
  		if name != "" {
    		query += " AND"
  		} else {
    		query += " WHERE"
  		}
 		query += " age='" + ages[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": "Something Wrong",})
	}

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

func InsertUser(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	defer db.Close()

	err := c.Request().ParseForm()
	if err != nil {
		log.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": "Something Wrong",})
	}

	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	country := c.FormValue("country")

	_, errQuery := db.Exec("INSERT INTO users (name, age, address, country) VALUES (?, ?, ?, ?)", name, age, address, country)
	
	var response Response
	if errQuery != nil {
		response.Status = 400
		response.Message = "Error"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Status = 200
		response.Message = "Success"
		return c.JSON(http.StatusOK, response)
	}
}