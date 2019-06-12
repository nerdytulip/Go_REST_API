package mux

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct{
	 Title string `json:"Title"`
	 Desc string `json:"desc"`
	 Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter,r *http.Request){
	articles := Articles{
		Article{Title:"test title",Desc:"test description",Content:"hello world"},
	}

	fmt.Println("Endpoint hi:all articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"homepage endpoint hit")
}

func handleRequest(){

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":8081",myRouter))
}

func main(){
	handleRequest()
}