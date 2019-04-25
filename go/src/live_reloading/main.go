package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Testing the live reloading...")
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Starting server on port 3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
