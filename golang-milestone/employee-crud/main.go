package main

import (
	"employee-crud/database"
	"employee-crud/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting apllication...")

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.RoutesMap(server, db)

	http.ListenAndServe(":8080", server)

}
