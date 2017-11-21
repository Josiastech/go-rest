package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	fmt.Println("Hello World!")
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", router))
}
