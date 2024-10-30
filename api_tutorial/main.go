package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Topic string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{
			Title:"Sample Title",
			Desc:"Sample description here to show some information",
			Topic:"Learning GO Rest API"},
	}

	fmt.Println("Endpoint at All Articles")
	json.NewEncoder(w).Encode(articles)
}

func createArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Create Articles Endpoint Testing")
}

const PORT int64 = 3000
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Server Homepage")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles",allArticles).Methods("GET")
	myRouter.HandleFunc("/articles",createArticles).Methods("POST")
	fmt.Printf("Server is running at port %d\n",PORT)
	log.Fatal(http.ListenAndServe(":3000",myRouter))
}

func main(){
	handleRequests()
}