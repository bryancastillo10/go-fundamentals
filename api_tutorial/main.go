package main

import (
	"fmt"
	"net/http"
)

func main(){
	const PORT int64 = 3000
	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"All comments can be returned")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id")
		fmt.Fprintf(w,"A path with id %s single comment is returned",id)
	})

		mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Post a new comment")
	})

	fmt.Printf("Server is running on port %d\n",PORT)
	if err := http.ListenAndServe("localhost:3000",mux); err != nil {
		fmt.Println(err.Error())
	}
}