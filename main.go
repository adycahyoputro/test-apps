package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/Books", helloHandler)
	
	log.Println("listening on port: 8081 always listening always undestanding")
	http.ListenAndServe(":8081", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid request method")
		return
	}

	// fmt.Fprintf(w, "hello world")
	fmt.Fprint(w, "a")
}