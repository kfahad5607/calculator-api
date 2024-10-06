package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type operands struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type numbers []float64

type response struct {
	Result float64 `json:"result"`
}

type invalidInputResponse[T any] struct {
	Error string `json:"error"`
	ExpectedInput T
}

type responseError struct {
	Error string `json:"error"`
}

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inputOperands operands
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&inputOperands)
	if err != nil{
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)

		errResponse := invalidInputResponse[operands]{Error: "Invalid input", ExpectedInput: operands{ Number1: 20, Number2: 5}}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	
	response := response{Result: inputOperands.Number1 + inputOperands.Number2}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}


func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inputOperands operands
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&inputOperands)
	if err != nil{
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)

		errResponse := invalidInputResponse[operands]{Error: "Invalid input", ExpectedInput: operands{ Number1: 20, Number2: 5}}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	
	response := response{Result: inputOperands.Number1 - inputOperands.Number2}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func HandleMultiply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inputOperands operands
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&inputOperands)
	if err != nil{
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)

		errResponse := invalidInputResponse[operands]{Error: "Invalid input", ExpectedInput: operands{ Number1: 20, Number2: 5}}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	
	response := response{Result: inputOperands.Number1 * inputOperands.Number2}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func HandleDivide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inputOperands operands
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&inputOperands)
	if err != nil{
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)

		errResponse := invalidInputResponse[operands]{Error: "Invalid input", ExpectedInput: operands{ Number1: 20, Number2: 5}}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if inputOperands.Number2 == 0 {
		w.WriteHeader(http.StatusBadRequest)
		errResponse := responseError{Error: "Denominator cannot be zero."}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := response{Result: inputOperands.Number1 / inputOperands.Number2}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func HandleSum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inputNumbers numbers
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	
	err := decoder.Decode(&inputNumbers)
	if err != nil{
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)

		errResponse := invalidInputResponse[numbers]{Error: "Invalid input", ExpectedInput: numbers{ 10, 20, 30}}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	
	w.WriteHeader(http.StatusOK)

	var result float64
	for _, num := range inputNumbers{
		result += num
	}

	response := response{Result: result} 
	json.NewEncoder(w).Encode(response)
}