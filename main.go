package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person blablabla
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var arr = make([]Person, 12)

func handlePeoples(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(arr)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/peoples", handlePeoples).Methods("GET")
	log.Println("Server is Running...")
	if err := http.ListenAndServe(":8081", r); err != nil {
		panic(err)
	}
}

func main() {
	for index := 0; index < len(arr); index++ {
		arr[index] = Person{index, "test"}
	}
	handleRequests()
}
