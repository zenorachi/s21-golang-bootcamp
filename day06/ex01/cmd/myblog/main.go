package main

import (
	"day06/ex01/internal/handlers"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./ui/img"))
	http.Handle("/img/", http.StripPrefix("/img/", fs))

	http.HandleFunc("/", handlers.HandleDefault)

	http.HandleFunc("/admin", handlers.HandleAdmin)

	log.Println("Server started ...")
	log.Fatalln(http.ListenAndServe("localhost:7777", nil))
}
