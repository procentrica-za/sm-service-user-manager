package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var config Config

//create initialisation functions
func init() {
	config = CreateConfig()
	fmt.Println("Config file has loaded")
	fmt.Printf("CrudHost: %v\n", config.CRUDHost)
	fmt.Printf("CrudPort: %v\n", config.CRUDPort)
	fmt.Printf("EmailHost: %v\n", config.EMAILHost)
	fmt.Printf("EmailPort: %v\n", config.EMAILPort)
	fmt.Printf("UsermanagerPort: %v\n", config.USERMANAGERPort)
}

//create config functions
func CreateConfig() Config {
	conf := Config{
		CRUDHost:        os.Getenv("CRUD_Host"),
		CRUDPort:        os.Getenv("CRUD_Port"),
		EMAILHost:       os.Getenv("EMAIL_Host"),
		EMAILPort:       os.Getenv("EMAIL_Port"),
		USERMANAGERPort: os.Getenv("USER_MANAGER_PORT"),
	}
	return conf
}
func main() {
	server := Server{
		router: mux.NewRouter(),
	}
	//Set up routes for server
	server.routes()
	handler := removeTrailingSlash(server.router)
	fmt.Print("starting server on port " + config.USERMANAGERPort + "\n")
	log.Fatal(http.ListenAndServe(":"+config.USERMANAGERPort, handler))
}
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
