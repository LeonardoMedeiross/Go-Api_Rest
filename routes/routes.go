package routes

import (
	"log"
	"net/http"

	"github.com/LeonardoMedeiross/Go-Api_Rest/controlles"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controlles.Home)
	r.HandleFunc("/api/personalidades", controlles.TodasAsPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controlles.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
