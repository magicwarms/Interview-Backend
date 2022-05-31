package main

import (
	"basic-api/infrastructures"
	"fmt"
	"log"
	"net/http"
)
func handleRequests() {
	// all app routes here
	appRoute := infrastructures.Dispatch()
	
	fmt.Println("⚡️ [Basic-API] - IS RUNNING ON PORT - 9000⚡️")
	log.Fatal(http.ListenAndServe(":9000", appRoute))
}

func main() {
	handleRequests()
}