package domain

type User struct {
	Id     int    `json:"id"`
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Fname  string `json:"fname"`
}

type UserRequest struct {
	Id     int    `json:"id"`
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Fname  string `json:"fname"`
}

type UserResponse struct {
	Data    []User `json:"data"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}
