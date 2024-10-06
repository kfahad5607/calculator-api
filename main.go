package main

import (
	"log"
	"net/http"
	"github.com/kfahad5607/calculator-api/handlers"
	"github.com/kfahad5607/calculator-api/middlewares"
)


func main(){
	router := http.NewServeMux()

	router.HandleFunc("POST /add", handlers.HandleAdd)
	router.HandleFunc("POST /subtract", handlers.HandleSubtract)
	router.HandleFunc("POST /multiply", handlers.HandleMultiply)
	router.HandleFunc("POST /divide", handlers.HandleDivide)
	router.HandleFunc("POST /sum", handlers.HandleSum)

	err := http.ListenAndServe(":3000", middlewares.Logger(router))
	if err != nil {
		log.Fatal(err)
	}
}