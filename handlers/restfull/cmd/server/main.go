package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/go-folder-structure/handlers/restfull"
)

func main() {
	http.HandleFunc("/", restfull.RestFull)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
