package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/mux"
)

var counter uint64

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{}

func main() {
	items = append(items, Item{Id: "1", Name: "books"})

	atomic.AddUint64(&counter, 0)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItemByID).Methods("GET")
	router.HandleFunc("/items", AddItems).Methods("POST")
	router.HandleFunc("/items", DeleteAllItems).Methods("DELETE")
	router.HandleFunc("/items/{id}", DeleteItemByID).Methods("DELETE")

	fmt.Println("GO Rest API server for items is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	fmt.Println("Endpoint hit: GetItems")
	json.NewEncoder(w).Encode(items)
}

func getItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("CONTENT-TYPE", "application/json")

	vars := mux.Vars(r)

	id := vars["id"]
	fmt.Println("Endpoint hit: getItemByID")
	for _, item := range items {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func AddItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("CONTENT-TYPE", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)
	items = []Item{}
	json.Unmarshal(reqBody, &items)

	fmt.Println("Endpoint hit: AddItems")
	json.NewEncoder(w).Encode(items)

}

func DeleteAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("CONTENT-TYPE", "application/json")

	items = []Item{}
	fmt.Println("Endpoint hit: DeleteAllItems")
	json.NewEncoder(w).Encode(items)
}

func DeleteItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("CONTENT-TYPE", "application/json")

	vars := mux.Vars(r)

	id := vars["id"]
	fmt.Println("Endpoint hit: DeleteItemByID")
	for index, item := range items {
		if item.Id == id {
			items = append(items[:index], items[index+1:]...)
		}
	}

}
