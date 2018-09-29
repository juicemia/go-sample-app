package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("booting sample app...")

	http.HandleFunc("/", handleRoot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		wr.WriteHeader(http.StatusNoContent)
		return
	default:
		wr.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
