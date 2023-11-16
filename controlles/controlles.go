package controlles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeonardoMedeiross/Go-Api_Rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasAsPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
