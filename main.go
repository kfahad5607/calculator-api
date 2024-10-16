package main

import (
	"log"
	"net/http"

	"github.com/kfahad5607/calculator-api/api/handlers"
	"github.com/kfahad5607/calculator-api/api/middlewares"
)

func main(){
	router := http.NewServeMux()

	router.HandleFunc("POST /add", handlers.HandleAdd)
	router.HandleFunc("POST /subtract", handlers.HandleSubtract)
	router.HandleFunc("POST /multiply", handlers.HandleMultiply)
	router.HandleFunc("POST /divide", handlers.HandleDivide)
	router.HandleFunc("POST /sum", handlers.HandleSum)

	// err := http.ListenAndServe(":3000", api.Logger(api.RateLimiter(router)))
	// err := http.ListenAndServe(":3000", api.CreateMiddlewareStack(router, api.Logger, api.RateLimiter))
	finalMiddleware := middlewares.CreateMiddlewareStack(middlewares.Logger, middlewares.RateLimiter)
	err := http.ListenAndServe(":3000", finalMiddleware(router))
	if err != nil {
		log.Fatal(err)
	}
}