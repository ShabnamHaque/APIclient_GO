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
	Desc  string `json:"desc"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "TEST", Desc: "First article"},
		Article{Title: "TEST", Desc: "second article"},
	}
	fmt.Println("Endpoint Hit : All articles endpoint")
	json.NewEncoder(w).Encode(articles)
	fmt.Fprintf(w, "All articles hit")

}
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage hit")
}
func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//		StrictSlash(true)	 /article and /article/ same
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/article", allArticles).Methods("GET")
	myRouter.HandleFunc("/article", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", myRouter))
}
func main() {
	fmt.Println("Starting Server at port 9090")
	handleRequests()
}
