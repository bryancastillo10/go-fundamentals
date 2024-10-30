package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

const PORT int64 = 3000
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Server Homepage")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles",allArticles)
	fmt.Printf("Server is running at port %d\n",PORT)
	log.Fatal(http.ListenAndServe(":3000",nil))
}

func main(){
	handleRequests()
}