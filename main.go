package main

import (
	"fmt"

	"github.com/LeonardoMedeiross/Go-Api_Rest/database"
	"github.com/LeonardoMedeiross/Go-Api_Rest/models"
	"github.com/LeonardoMedeiross/Go-Api_Rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome1", Historia: "historia1"},
		{Id: 2, Nome: "nome2", Historia: "historia2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("iniciando o servidor com Rest com GO")
	routes.HandleResquest()
}
