package main

import (
	"fmt"

	"github.com/LeonardoMedeiross/Go-Api_Rest/models"
	"github.com/LeonardoMedeiross/Go-Api_Rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "nome1", Historia: "historia1"},
		{Nome: "nome2", Historia: "historia2"},
	}

	fmt.Println("iniciando o servidor com Rest com GO")
	routes.HandleResquest()
}
