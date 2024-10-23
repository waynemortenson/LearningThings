package main

import(
	"log"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, world!!"))
	log.Print("Serving data")
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting server on :8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}