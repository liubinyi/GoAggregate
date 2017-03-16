package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func fetchHello(w http.ResponseWriter, r *http.Request) {

}

func fetchWorld(w http.ResponseWriter, r *http.Request) {

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/awesomeapi/{name}", fetchHello).Methods("GET")
	router.HandleFunc("/awesomego/{id}", fetchWorld).Methods("GET")
	http.ListenAndServe(":8080", router)
	//":"+os.Getenv("PORT")
}
