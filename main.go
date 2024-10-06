package main

import (
	"fmt"
	"log"
	"net/http"

	models "GolangAPIProject/model"
	"GolangAPIProject/router"

	"github.com/rs/cors"
)

func main() {
	models.ConnectDB() // Connect to the database

	r := router.Router() // Get the router instance

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins or specify your front-end URL
		AllowCredentials: true,
	})

	fmt.Println("Starting server on the port 4000...")
	log.Fatal(http.ListenAndServe(":4000", c.Handler(r))) // Start the server
}
