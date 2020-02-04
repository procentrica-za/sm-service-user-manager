package main

import "github.com/gorilla/mux"

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}
type UpdateUser struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}
type DeleteUser struct {
	UserID string `json:"id"`
}
type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type GetUser struct {
	UserID string `json:"id"`
}
type ForgetPassword struct {
	Email string `json:"email"`
}

type UpdateUserResult struct {
	UserUpdated bool   `json:"userupdated"`
	Message     string `json:"message"`
}
type LogoutUser struct {
	UserID string `json:"id"`
}
type RegisterUserResult struct {
	UserCreated string `json:"usercreated"`
	Username    string `json:"username"`
	UserID      string `json:"id"`
	Message     string `json:"message"`
}
type GetUserResult struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}
type DeleteUserResult struct {
	UserDeleted bool   `json:"userdeleted"`
	UserID      string `json:"id"`
	Message     string `json:"message"`
}
type LoginUserResult struct {
	UserLoggedIn bool   `json:"userloggedin"`
	UserID       string `json:"id"`
	Message      string `json:"message"`
}
type ForgetPasswordResult struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LogoutUserResult struct {
	Username string `json:"username"`
}
type Server struct {
	router *mux.Router
}
type Config struct {
	CRUDHost string
	CRUDPort string
}
