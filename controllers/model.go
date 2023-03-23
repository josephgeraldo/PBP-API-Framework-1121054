package controllers

//User is the model for user
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Country string `json:"country"`
}

//UsersResponse is the response for get all users
type UsersResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

//ErrorResponse is the response for error
type ErrorResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}

//Response is the response for get, post, put
type Response struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data    User `json:"data"`
}

//ResponseDelete is the response for delete
type ResponseDelete struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	ID	    int `json:"id"`
}