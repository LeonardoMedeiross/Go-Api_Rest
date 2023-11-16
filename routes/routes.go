package routes

import (
	"log"
	"net/http"

	"github.com/LeonardoMedeiross/Go-Api_Rest/controlles"
)

func HandleResquest() {
	http.HandleFunc("/", controlles.Home)
	http.HandleFunc("/api/personalidades", controlles.TodasAsPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
