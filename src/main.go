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

func init() {
	config = CreateConfig()
	fmt.Println("Config file has loaded")
	fmt.Printf("CrudHost: %v\n", config.CRUDHost)
	fmt.Printf("CrudPort: %v\n", config.CRUDPort)
}
func CreateConfig() Config {
	conf := Config{
		CRUDHost: os.Getenv("CRUD_Host"),
		CRUDPort: os.Getenv("CRUD_Port"),
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
	fmt.Print("starting server on port 8888\n")
	log.Fatal(http.ListenAndServe(":8888", handler))
}
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
