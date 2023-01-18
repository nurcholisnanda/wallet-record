package main

import (
	"log"
	"os"

	"github.com/nurcholisnanda/wallet-record/controllers"
	"github.com/nurcholisnanda/wallet-record/repositories"
	"github.com/nurcholisnanda/wallet-record/routes"
	"github.com/nurcholisnanda/wallet-record/services"
)

func main() {
	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	// Initialize the router
	set := routes.NewRouter(controllers.NewController(services.NewService(repositories.NewRepository())))
	r := set.SetupRouter()
	// Start the server
	r.Run(":" + port)
}
