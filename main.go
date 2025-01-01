package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type DbBook struct {
	Id     int
	Title  string
	Author string
}

var cfg = []DbBook{
	{Id: 1, Title: "buku sejarah", Author: "budi"},
	{Id: 2, Title: "buku fisika", Author: "rara"},
	{Id: 3, Title: "buku ppkn", Author: "nana"},
}

func main() {
	http.HandleFunc("/Books", GetBookList)
	http.HandleFunc("/Book1", GetBookById)
	// http.HandleFunc("/Book2", InsertBook)
	// http.HandleFunc("/Book3", UpdateBook)
	// http.HandleFunc("/Book4", DeleteBook)
	log.Println("listening on port: 8080 always listening always undestanding")
	http.ListenAndServe(":8080", nil)

}

func GetBookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// bookList := repo.Books
		var result, err = json.Marshal(cfg)
		fmt.Println(cfg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
		// json.NewEncoder(w).Encode(cfg)
	}
	http.Error(w, "status bad request", http.StatusBadRequest)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id := r.FormValue("id")
		idInt, errStr := strconv.Atoi(id)
		if errStr != nil {
			fmt.Println(errStr)
		}
		var result []byte
		var err error
		bookList := cfg
		for _, value := range bookList {

			if value.Id == idInt {
				result, err = json.Marshal(value)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	http.Error(w, "status bad request", http.StatusBadRequest)
}