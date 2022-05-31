package controllers

import (
	"basic-api/config"
	"basic-api/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func InitialRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(config.AppResponse("Hello World"))
}

func BackendQuestionOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	getData, err := services.QuestionOne()
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(config.AppResponse(getData))
}

func BackendQuestionTwo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	getData, err := services.QuestionTwo()
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(config.AppResponse(getData))
}

func BackendQuestionThree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	getData, err := services.QuestionThree()
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(config.AppResponse(getData))
}
