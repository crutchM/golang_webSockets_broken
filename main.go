package main

import (
	"back_edu/model"
	"log"
	"net/http"
)

func main() {
	server := model.NewServer("/entry")
	go server.Listen()
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	log.Fatal(http.ListenAndServe(":80", nil))
}
