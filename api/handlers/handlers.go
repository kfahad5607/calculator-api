package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Name string `json:"name"`
}

type LoginCreds struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type operands struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type numbers []float64

type response[T any] struct {
	Result T `json:"result"`
}
type invalidInputResponse[T any] struct {
	Error string `json:"error"`
	ExpectedInput T
}

type ResponseError struct {
	Error string `json:"error"`
}

var Users map[string]User = map[string]User{
	"user.one@example.com": {
		Id: 1,
		Email: "user.one@example.com",
		Name: "User One", 
		PasswordHash: "$2a$10$S2nth4lumoXlkblPfnLcgu28xqieTzOBFaHT1wg4uNm/nHf9HGMCS",
	},
	"user.two@example.com": {
		Id: 2,
		Email: "user.two@example.com", 
		Name: "User Two", 
		PasswordHash: "$2a$10$xiM/frM7QsDR17mfMpvGx.J1Qji2/vEVXbUY41kXa3w06Wq01nr1K",
	},
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var loginCreds LoginCreds
	decoder.Decode(&loginCreds)

	user, exists := Users[loginCreds.Email]
	encoder := json.NewEncoder(w)
	if !exists {
		errResponse := ResponseError{Error: "Invalid credentials!"}
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(errResponse)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginCreds.Password))
	if err != nil{
		errResponse := ResponseError{Error: "Invalid credentials!"}
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(errResponse)
		return
	}

	ACCESS_TOKEN_EXPIRE_IN := os.Getenv("ACCESS_TOKEN_EXPIRE_IN")
	expireIn, err :=strconv.Atoi(ACCESS_TOKEN_EXPIRE_IN)
	if err != nil {
		fmt.Println(err)
	}

	issuedAt := time.Now().UTC().Unix()
	expiredAt := time.Now().UTC().Add(time.Minute * time.Duration(expireIn)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": issuedAt,
		"exp": expiredAt,
		"id": user.Id,
		"name": user.Name,
		"email": user.Email,
	})

	ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")
	tokenString, _err := token.SignedString([]byte(ACCESS_TOKEN_SECRET))
	if _err != nil{
		fmt.Println(_err)
	}

	response := response[string]{Result: tokenString}
	encoder.Encode(response)
	w.WriteHeader(http.StatusOK)
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
	
	response := response[float64]{Result: inputOperands.Number1 + inputOperands.Number2}

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
	
	response := response[float64]{Result: inputOperands.Number1 - inputOperands.Number2}

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
	
	response := response[float64]{Result: inputOperands.Number1 * inputOperands.Number2}

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
		errResponse := ResponseError{Error: "Denominator cannot be zero."}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := response[float64]{Result: inputOperands.Number1 / inputOperands.Number2}

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

	response := response[float64]{Result: result} 
	json.NewEncoder(w).Encode(response)
}