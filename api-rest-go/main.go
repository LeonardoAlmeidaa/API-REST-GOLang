package main

import (
	"api-rest/database"
	"api-rest/models"
	"api-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Leonardo", Historia: "História do Léo"},
		{Id: 2, Nome: "Eduardo", Historia: "História do Edu"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o Servidor REST GO")
	routes.HandleRequest()
}
