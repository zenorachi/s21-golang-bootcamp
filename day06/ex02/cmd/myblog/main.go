package main

import (
	"day06/ex02/internal/handlers"
	"log"
	"net/http"
)

const PORT = "8888"

func main() {
	img := http.FileServer(http.Dir("./ui/img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))

	md := http.FileServer(http.Dir("./ui/md/"))
	http.Handle("/md/", http.StripPrefix("/md/", md))

	http.HandleFunc("/", handlers.HandleDefault)

	http.HandleFunc("/admin", handlers.HandleAdmin)

	log.Println("Server started at PORT:", PORT+"...")
	log.Fatalln(http.ListenAndServe("localhost:"+PORT, nil))
}
