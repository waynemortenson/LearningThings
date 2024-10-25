package main

import(
	"fmt"
	"log"
	"net/http"
	"strconv"
)


func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, world!!"))
	log.Print("Serving data")
}

func snippetView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	log.Print("Serving snippet view")
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	log.Print("Serving snippet create")
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/items/{item}", Item)
	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}