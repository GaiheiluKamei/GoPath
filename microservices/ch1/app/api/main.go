package main

import (
	"github.com/GaiheiluKamei/books/microservices/ch1/app/api/handlers"
	"github.com/GaiheiluKamei/books/microservices/ch1/app/api/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewRepository("mongodb://localhost:27017", "packt", "timeZones")
	defer repo.Close()

	h := handlers.Handlers{Repo:repo}

	r := mux.NewRouter()
	r.HandleFunc("/timeZones", h.All).Methods("GET")
	r.HandleFunc("/timeZones/{timeZone}", h.GetByTZ).Methods("GET")

	r.HandleFunc("/timeZones", h.Insert).Methods("POST")
	r.HandleFunc("/timeZones/{timeZone}", h.Delete).Methods("DELETE")
	r.HandleFunc("/timeZones/{timeZone}", h.Update).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8081", r))
}
