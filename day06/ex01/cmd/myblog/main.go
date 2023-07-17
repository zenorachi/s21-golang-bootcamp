package main

import (
	"day06/ex01/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/admin", handlers.HandleAdmin)
	http.HandleFunc("/", handlers.HandleDefault)
	log.Println("Server started ...")
	log.Fatalln(http.ListenAndServe("localhost:8888", nil))
}
