package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Vehicle struct {
	Id   string
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var v Vehicle
	if r.Body == nil {
		http.Error(w, "Invalid request", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("----Server 1 --- Host: localhost:8080----")
	fmt.Println("I received payload")
	fmt.Println(v.Id)
	fmt.Println(v.Name)
	fmt.Println("End PayLoad")
	fmt.Println("------------------------------------------")
}

func main() {
	http.HandleFunc("/handleme", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
