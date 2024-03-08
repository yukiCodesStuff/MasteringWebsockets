// This will be running our server
package main

import (
	"net/http"
	"log"
)

func main() {
	setupAPI()

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func setupAPI() {

	manager := NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
}