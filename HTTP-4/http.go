package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type article struct {
	Title   string `json:title`
	Desc    string `json:Desc`
	Content string `json:String`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnallarticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Return All Articles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/article", returnallarticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

var Articles []article

func main() {
	Articles = []article{
		article{Title: "human", Desc: "Living Being", Content: "Alive"},
		article{Title: "Dog", Desc: "Animal", Content: "Caring"},
	}
	handleRequests()
}
