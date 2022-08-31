package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello interwebs</h1>")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h2> About Time I learnt GO </h2>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/About", about)

	http.ListenAndServe(":3000", nil)
	fmt.Println("Server running")
}