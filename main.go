package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{
			Id:       1,
			Nome:     "Nome 1",
			Historia: "Historia",
		},
		{
			Id:       2,
			Nome:     "Nome 2",
			Historia: "Historia 2",
		},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando Servidor...")
	routes.HandleRequests()
}
