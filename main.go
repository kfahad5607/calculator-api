package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kfahad5607/calculator-api/api/handlers"
	"github.com/kfahad5607/calculator-api/api/middlewares"
)

func main(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	router := http.NewServeMux()
	protectedRoute := http.NewServeMux()
	
	protectedRoute.HandleFunc("POST /add", handlers.HandleAdd)
	protectedRoute.HandleFunc("POST /subtract", handlers.HandleSubtract)
	protectedRoute.HandleFunc("POST /multiply", handlers.HandleMultiply)
	protectedRoute.HandleFunc("POST /divide", handlers.HandleDivide)
	protectedRoute.HandleFunc("POST /sum", handlers.HandleSum)
	
	router.HandleFunc("POST /login", handlers.Login)
	router.Handle("/", middlewares.CheckAuth(protectedRoute))

	finalMiddleware := middlewares.CreateMiddlewareStack(middlewares.Logger, middlewares.RateLimiter)
	err = http.ListenAndServe(":3000", finalMiddleware(router))
	if err != nil {
		log.Fatal(err)
	}
}