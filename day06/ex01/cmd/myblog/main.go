package main

import (
	"day06/ex01/internal/database"
	"fmt"
	"log"
)

func main() {
	articles, err := database.GetArticles()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(articles)
	//http.HandleFunc("/admin", handlers.HandleAdmin)
	//log.Println("Server started ...")
	//log.Fatalln(http.ListenAndServe("localhost:8888", nil))
}
