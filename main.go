package main

import (
	"log"
	"net/http"
	"time"

	delivery "github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/delivery/http"
	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/downstream"
	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/service"
	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/config"
	"github.com/gorilla/mux"
)

func main() {
	// The main function where all the init happens and the servers starts up

	// build the downstream implementation
	downstream := downstream.NewSochain(config.API_BASE_URL)
	// build the service
	svc := service.NewBlockAPI(downstream)

	// build the handler
	r := mux.NewRouter()
	server := &http.Server{
		Handler:      r,
		Addr:         config.ADDRESS,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	delivery.NewController(r, svc)
	// start the sever

	log.Fatal(server.ListenAndServe())
}
