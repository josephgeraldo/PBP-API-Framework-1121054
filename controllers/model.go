package controllers

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Country string `json:"country"`
}

type UsersResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ErrorResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}
type Response struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data    User `json:"data"`
}

type ResponseDelete struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Id	    int `json:"id"`
}