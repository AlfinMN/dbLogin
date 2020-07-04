package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer(router *mux.Router) {
	host := "localhost"
	port := "4646"
	fmt.Println("succes connect to port : " + port)
	err := http.ListenAndServe(host+":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
