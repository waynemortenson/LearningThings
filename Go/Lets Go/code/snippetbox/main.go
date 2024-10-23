package main

import(
	"log"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, world!!"))
	log.Print("Serving data")
}

func snippetView(w http.ResponseWriter, r *http.Request){
	log.Print("Serving snippet view")
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	log.Print("Serving snippet create")
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Starting server on :8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}