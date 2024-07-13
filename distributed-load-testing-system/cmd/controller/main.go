package main

import (
	"distributed-load-testing-system/internal/middleware"
	"distributed-load-testing-system/pkg/controller/handlers"
	"distributed-load-testing-system/pkg/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ram28873"
	dbname   = "dlts"
)

//To connect in Linux terminal sudo -u postgres psql

var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func main() {
	router := mux.NewRouter()

	//Middleware
	router.Use(middleware.LoggingMiddleware)

	postgres_conn_err := storage.InitDB(psqlInfo)
	if postgres_conn_err != nil {
		log.Println("Failed to Connect with Postgres Database 'dlts'")
	} else {
		log.Println("Connected to Postgres Database 'dlts'")
	}

	// Routes
	router.HandleFunc("/config", handlers.CreateConfig).Methods("POST")
	router.HandleFunc("/config/{id}", handlers.GetConfig).Methods("GET")
	router.HandleFunc("/config/{id}", handlers.UpdateConfig).Methods("PUT")
	router.HandleFunc("/config/{id}", handlers.DeleteConfig).Methods("DELETE")

	log.Println("Starting server on :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
