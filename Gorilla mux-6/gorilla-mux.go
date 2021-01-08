package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type article struct {
	ID      string `json:ID`
	Title   string `json:Title`
	Desc    string `json: Desc`
	Content string `json: Content`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to gorilla mux homepage")
	fmt.Println("Endpoint hit: Homepage")
}

func createnewarticle(w http.ResponseWriter, r *http.Request) {
	reqbody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Create new article")
	fmt.Println("Endpoint Hit: Create new article")
	var Article article
	json.Unmarshal(reqbody, &Article)
	Articles = append(Articles, Article)
	json.NewEncoder(w).Encode(Articles)
}

func updatearticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["ID"]

	reqbody, _ := ioutil.ReadAll(r.Body)
	var Article article
	json.Unmarshal(reqbody, &Article)

	for _, article := range Articles {
		if article.ID == keys {
			i, err := strconv.Atoi(keys)
			if err == nil {
				fmt.Println(err)
			} else {
				Articles = append(Articles[:i], Articles[i+1:]...)
				Articles = append(Articles, article)
				json.NewEncoder(w).Encode(Articles)
			}
		}
	}
}

func deletearticleid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["ID"]
	fmt.Println("Endpointhit: Delete Article by ID")

	for _, article := range Articles {
		if article.ID == keys {
			i, err := strconv.Atoi(keys)
			if err != nil {
				fmt.Println(w, err)
			}
			Articles = append(Articles[:i], Articles[i+1:]...)
			fmt.Fprintf(w, keys, "Aricle deleted")
			json.NewEncoder(w).Encode(Articles)
		}
	}
}

func returnarticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: Return Article")
	json.NewEncoder(w).Encode(Articles)
}

func returnonearticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["ID"]
	fmt.Println("Endpoint Hit: Return Single article")

	for _, Article := range Articles {
		if Article.ID == keys {
			json.NewEncoder(w).Encode(Article)
		}
	}
}

func handlerequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homepage)
	router.HandleFunc("/article", createnewarticle).Methods("POST")
	router.HandleFunc("/article/{ID}", returnonearticle).Methods("GET")
	router.HandleFunc("/article", returnarticle).Methods("GET")
	router.HandleFunc("/article/{ID}", updatearticle).Methods("PUT")
	router.HandleFunc("/delete/{ID}", deletearticleid).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10001", router))
}

var Articles []article

func main() {
	fmt.Println("Mux Router")
	Articles = []article{
		article{ID: "1", Title: "boss", Desc: "Movie", Content: "Sivaji"},
		article{ID: "2", Title: "100", Desc: "Police", Content: "Emergency"},
	}
	handlerequest()
}
