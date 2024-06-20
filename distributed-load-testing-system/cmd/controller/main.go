package main

import (
	"distributed-load-testing-system/internal/middleware"
	"distributed-load-testing-system/pkg/controller/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//Middleware
	router.Use(middleware.LoggingMiddleware)

	// Routes
	router.HandleFunc("/config", handlers.CreateConfig).Methods("POST")
	router.HandleFunc("/config/{id}", handlers.GetConfig).Methods("GET")
	router.HandleFunc("/config/{id}", handlers.UpdateConfig).Methods("PUT")
	router.HandleFunc("/config/{id}", handlers.DeleteConfig).Methods("DELETE")

	log.Println("Starting server on :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
