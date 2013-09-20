package main

import (
	"net/http"
	"log"

	"./handler"
)


func main() {
	http.HandleFunc("/hello", handler.HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


