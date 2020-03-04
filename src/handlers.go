package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) handleregisteruser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		regUser := RegisterUser{}
		err := json.NewDecoder(r.Body).Decode(&regUser)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Improper registration details provided")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(regUser)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/user", "application/json", bytes.NewBuffer(requestByte))

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to register")
			return
		}
		if req.StatusCode != 200 {
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to process registration")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)

			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed..."+bodyString)
			fmt.Println("Request to DB can't be completed..." + bodyString)
			return
		}
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Registration is not able to be completed by internal error")
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var registerResponse RegisterUserResult

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)

		err = decoder.Decode(&registerResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding registration response")
			return
		}
		js, jserr := json.Marshal(registerResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the response to register user")
			return
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
func (s *Server) handleupdateuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		updateUser := UpdateUser{}
		err := json.NewDecoder(r.Body).Decode(&updateUser)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		client := &http.Client{}

		//create byte array from JSON payload
		requestByte, _ := json.Marshal(updateUser)

		//put to crud service
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/user", bytes.NewBuffer(requestByte))
		if err != nil {
			fmt.Fprint(w, err.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to update user")
			return
		}

		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		//close the request
		defer resp.Body.Close()

		//create new response struct
		var updateResponse UpdateUserResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&updateResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(updateResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the response to update user")
			return
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
func (s *Server) handledeleteuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get id of user for deletion
		userid := r.URL.Query().Get("id")
		client := &http.Client{}

		//send delete request to CRUD service
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/user?id="+userid, nil)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to delete a user")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		//close the request
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println("An error has occured whilst sending user ID for deletion")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed..."+bodyString)
			fmt.Println("Request to DB can't be completed..." + bodyString)
			return
		}

		//create new response struct
		var deleteResponse DeleteUserResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(deleteResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the response to delete a user")
			return
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
func (s *Server) handleloginuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get username and password for login
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")
		if username == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "No username provided in URL")
			fmt.Println("A username has not been provided in URL")
			return
		}
		if password == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "No password provided in URL")
			fmt.Println("A password has not been provided in URL")
			return
		}

		//get from CRUD service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/userlogin?username=" + username + "&password=" + password)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to login user")
			return
		}
		if req.StatusCode != 200 {
			fmt.Println("Request to DB can't be completed to login user")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Database error occured upon retrieval"+bodyString)
			fmt.Println("Database error occured upon retrieval" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var loginResponse LoginUserResult
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&loginResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Unable to decode login response")
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(loginResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the response to logging in a user")
			return
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get userID from url
		userid := r.URL.Query().Get("id")
		if userid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "UserID not properly provided in URL")
			fmt.Println("UserID not properly provided in URL")
			return
		}

		//get userID from crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/user?id=" + userid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve user information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(500)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get user data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get user data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var getResponse GetUserResult
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("An internal error has occured whilst trying to decode the get user response")
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(getResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the response to get user")
			return
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleforgotpassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle forgot password Has Been Called in the Cred manager and recieved the email from URL")
		//Get Email Address from URL

		useremail := r.URL.Query().Get("email")

		//Check if no Email address was provided in the URL
		if useremail == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Email address not properly provided in URL")
			fmt.Println("Email address not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/forgotpassword?email=" + useremail)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve password reset information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get advertisement data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get advertisement data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var forgotpasswordresult ForgotPasswordResult

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&forgotpasswordresult)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get Advertisement response ")
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(forgotpasswordresult)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//Take CRUD response and compile a new response for email service
		var forgotPasswordEmail ForgotPasswordEmail
		forgotPasswordEmail.ToEmail = forgotpasswordresult.Email
		forgotPasswordEmail.Subject = "Forgot Password"
		forgotPasswordEmail.Password = "Your new password as requested is: " + forgotpasswordresult.Password
		forgotPasswordEmail.Message = forgotpasswordresult.Message

		//create byte array from JSON payload
		requestByte1, _ := json.Marshal(forgotPasswordEmail)

		//send CRUD response to email service
		req1, respErr1 := http.Post("http://"+config.EMAILHost+":"+config.EMAILPort+"/forgotpassword", "application/json", bytes.NewBuffer(requestByte1))

		fmt.Println("Sent to email service")
		//check for response error of 500
		if respErr1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr1.Error())
			fmt.Println("Error received from email service->" + respErr1.Error())
			return
		}
		if req1.StatusCode != 200 {
			fmt.Fprint(w, "Request to Email Service can't be completed...")
			fmt.Println("Unable to process password reset")
		}
		if req1.StatusCode == 500 {
			w.WriteHeader(500)

			bodyBytes1, err1 := ioutil.ReadAll(req1.Body)
			if err1 != nil {
				log.Fatal(err1)
			}
			bodyString := string(bodyBytes1)
			fmt.Fprintf(w, "Request to Emal Service can't be completed..."+bodyString)
			fmt.Println("Request to Email Service can't be completed..." + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for front-end user
		var emailResponse EmailResult
		emailResponse.Message = forgotpasswordresult.Message

		//convert struct back to JSON
		js, jserr3 := json.Marshal(emailResponse)
		if jserr3 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr3.Error())
			fmt.Println("Error occured when trying to marshal the response to get user")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdatepassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		updatePassword := UpdatePassword{}
		err := json.NewDecoder(r.Body).Decode(&updatePassword)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		client := &http.Client{}

		//create byte array from JSON payload
		requestByte, _ := json.Marshal(updatePassword)

		//put to crud service
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/userpassword", bytes.NewBuffer(requestByte))
		if err != nil {
			fmt.Fprint(w, err.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to update user")
			return
		}

		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		//close the request
		defer resp.Body.Close()

		//create new response struct
		var passwordResponse UpdatePasswordResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&passwordResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(passwordResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the response to update user")
			return
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
