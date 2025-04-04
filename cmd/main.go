package main

import (
	"damt/config"
	"damt/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	log.Printf("🚀 Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
