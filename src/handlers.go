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
		regUser := RegisterUser{}
		err := json.NewDecoder(r.Body).Decode(&regUser)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Improper registration details provided")
			return
		}
		requestByte, _ := json.Marshal(regUser)
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/user", "application/json", bytes.NewBuffer(requestByte))
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to register")
			return
		}
		if req.StatusCode != 200 {
			fmt.Println(w, "Request to DB can't be completed...")
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
			fmt.Println(w, "Request to DB can't be completed..."+bodyString)
			return
		}
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Registration is not able to be completed - status 500")
			return
		}
		defer req.Body.Close()
		var registerResponse RegisterUserResult

		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&registerResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding registration result")
			return
		}
		fmt.Println(registerResponse.UserCreated)
		js, jserr := json.Marshal(registerResponse)
		fmt.Println(js)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println(w, "Error occured with the response of registration")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
func (s *Server) handleupdateuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateUser := UpdateUser{}
		err := json.NewDecoder(r.Body).Decode(&updateUser)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}

		client := &http.Client{}

		// Create request
		requestByte, _ := json.Marshal(updateUser)
		req, err := http.NewRequest("PUT", "http://"+config.CRUDHost+":"+config.CRUDPort+"/user", bytes.NewBuffer(requestByte))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error in communication with CRUD service endpoint for request to update user")
			return
		}

		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		var updateResponse UpdateUserResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&updateResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		js, jserr := json.Marshal(updateResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
func (s *Server) handledeleteuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.URL.Query().Get("id")
		client := &http.Client{}
		req, respErr := http.NewRequest("DELETE", "http://"+config.CRUDHost+":"+config.CRUDPort+"/user?id="+userid, nil)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to delete a user")
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println(w, "Request to DB can't be completed...")
		}
		if resp.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed..."+bodyString)
			fmt.Println(w, "Request to DB can't be completed..."+bodyString)
			return
		}
		var deleteResponse DeleteUserResult
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&deleteResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		js, jserr := json.Marshal(deleteResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
func (s *Server) handleloginuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")
		if username == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "No username")
			fmt.Println("No username provided")
			return
		}
		if password == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "No password")
			fmt.Println("No password provided")
			return
		}
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/userlogin?username=" + username + "&password=" + password)

		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to login user")
			return
		}
		if req.StatusCode != 200 {
			fmt.Println("Request to DB can't be completed...")
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
		defer req.Body.Close()
		var loginResponse LoginUserResult
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&loginResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Unable to decode login response")
			return
		}
		js, jserr := json.Marshal(loginResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("unable to marshal login response")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.URL.Query().Get("id")
		if userid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "UserID not properly provided in URL")
			fmt.Println("UserID not properly provided in URL")
			return
		}
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/user?id=" + userid)
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get user")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(500)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println(w, "Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed..."+bodyString)
			fmt.Println(w, "Request to DB can't be completed..."+bodyString)
			return
		}
		defer req.Body.Close()
		var getResponse GetUserResult
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		js, jserr := json.Marshal(getResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
func (s *Server) handleforgetpassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		forgetUser := ForgetPassword{}
		err := json.NewDecoder(r.Body).Decode(&forgetUser)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		requestByte, _ := json.Marshal(forgetUser)

		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/forgetpassword", "application/json", bytes.NewBuffer(requestByte))

		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request of forget password")
			return
		}
		if req.StatusCode != 200 {
			fmt.Println(w, "Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed..."+bodyString)
			fmt.Println(w, "Request to DB can't be completed..."+bodyString)
			return
		}
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			return
		}
		defer req.Body.Close()
		var forgetResponse ForgetPasswordResult
		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&forgetResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		js, jserr := json.Marshal(forgetResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
func (s *Server) handlelogoutuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logoutUser := LogoutUser{}
		err := json.NewDecoder(r.Body).Decode(&logoutUser)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		requestByte, _ := json.Marshal(logoutUser)

		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/userlogout", "application/json", bytes.NewBuffer(requestByte))

		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to logout user")
			return
		}
		if req.StatusCode != 200 {
			fmt.Println("Error in communication with CRUD service endpoint for request")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed..."+bodyString)
			fmt.Println(w, "Request to DB can't be completed..."+bodyString)
			return
		}
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			return
		}
		defer req.Body.Close()
		var logoutResponse LogoutUserResult
		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&logoutResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			return
		}
		js, jserr := json.Marshal(logoutResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
