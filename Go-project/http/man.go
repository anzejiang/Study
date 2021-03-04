package main

import (
	"fmt"
	"log"
	"net/http"
)











func sayhelloname(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "The is a web server.")
}

func main() {
	http.HandleFunc("/", sayhelloname)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
