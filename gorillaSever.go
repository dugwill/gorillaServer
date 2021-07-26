package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the HomePage!")
	log.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/printvars", printVars).Methods("POST")

	// Serve static files
	//myRouter.PathPrefix("/root_video0/").Handler(http.StripPrefix("/root_video0/", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// respondWithError Formats errors in JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON  writes the payload to the http response writer
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func printVars(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.URL.Query() {
		if k == "level" {
			fmt.Println(k, v)
		}
	}
}

func main() {

	go handleRequests()
	fmt.Println("Handel request started")

	for {
		time.Sleep(5 * time.Second)
		fmt.Println("still running")
	}
}
