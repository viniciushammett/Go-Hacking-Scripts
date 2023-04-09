package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWrite, r *http.Request){
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main(){
	http.HandleFunc("/hello", hello)
	http.ListenAndServer(":8000", nil)
}