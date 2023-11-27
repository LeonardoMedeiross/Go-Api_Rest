package routes

import (
	"log"
	"net/http"

	"github.com/LeonardoMedeiross/Go-Api_Rest/controlles"
	"github.com/LeonardoMedeiross/Go-Api_Rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controlles.Home)
	r.HandleFunc("/api/personalidades", controlles.TodasAsPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controlles.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controlles.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controlles.DeletandoUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controlles.EditarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
