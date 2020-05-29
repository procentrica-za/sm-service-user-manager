package main

import "github.com/gorilla/mux"

//create structs for JSON objects recieved and responses
type RegisterUser struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	InsitutionName string `json:"institutionname"`
}
type UpdateUser struct {
	UserID         string `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	InsitutionName string `json:"institutionname"`
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
	UserID         string `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	InsitutionName string `json:"institutionname"`
	AdsRemaining   string `json:"adsremaining"`
	Message        string `json:"message"`
	GotUser        bool   `json:"gotuser"`
}

type DeleteUserResult struct {
	UserDeleted bool   `json:"userdeleted"`
	UserID      string `json:"id"`
	Message     string `json:"message"`
}
type LoginUserResult struct {
	UserID       string `json:"id"`
	Username     string `json:"username"`
	UserLoggedIn bool   `json:"userloggedin"`
	Institution  string `json:"institution"`
	Message      string `json:"message"`
}
type ForgotPasswordResult struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Message  string `json:"message"`
}

type ForgotPasswordEmail struct {
	ToEmail  string `json:"toemail"`
	Subject  string `json:"subject"`
	Password string `json:"password"`
	Message  string `json:"message"`
}
type EmailResult struct {
	Message  string `json:"message"`
	Password string `json:"password"`
}
type LogoutUserResult struct {
	Username string `json:"username"`
}
type UpdatePassword struct {
	UserID          string `json:"id"`
	CurrentPassword string `json:"currentpassword"`
	Password        string `json:"password"`
}

type UpdatePasswordResult struct {
	PasswordUpdated bool   `json:"passwordupdated"`
	Message         string `json:"message"`
}

type InstitutionName struct {
	Institutionname string `json:"institutionname"`
}

type InstitutionNameList struct {
	Institutionnames []InstitutionName `json:"institutionnames"`
}

type ValidateOtp struct {
	UserID string `json:"userid"`
	Otp    string `json:"otp"`
}

type ValidateOtpResult struct {
	Validated bool   `json:"validated"`
	Message   string `'json:"message"`
}

type RequestOtpResult struct {
	Sent        bool   `json:"sent"`
	Message     string `'json:"message"`
	Phonenumber string `json:"phonenumber"`
	Otp         string `json:"otp"`
}

type SendText struct {
	Number  string `json:"number"`
	Message string `json:"message"`
}

type OtpResponse struct {
	Sent    bool   `json:"sent"`
	Message string `'json:"message"`
}

type Status struct {
	Isverified bool `json:"isverified"`
}

type PurchaseAdvertisement struct {
	ID      string `json:"id"`
	Ammount string `json:"ammount"`
}

type PurchaseAdvertisementResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

//touter service struct
type Server struct {
	router *mux.Router
}

type getPassword struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	GotUser  bool   `json:"gotuser"`
}

//config struct
type Config struct {
	CRUDHost        string
	CRUDPort        string
	USERMANAGERPort string
	EMAILHost       string
	EMAILPort       string
	TEXTHost        string
	TEXTPort        string
}
