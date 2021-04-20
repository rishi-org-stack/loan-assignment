package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	ser "github.com/rishi-org-stack/loan/server"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load(".env")

	router := ser.Route()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowCredentials: true,
	})
	srv := &http.Server{
		Handler: c.Handler(router),
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
