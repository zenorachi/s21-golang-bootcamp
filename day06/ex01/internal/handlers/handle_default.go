package handlers

import (
	"day06/ex01/internal/database"
	"log"
	"net/http"
)

//var tmplMain = template.Must(template.ParseFiles("./ui/html/articles.html"))

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/article1.html" {
		http.ServeFile(w, r, "./ui/html/article1.html")
	}
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./ui/html/articles.html")
	}
	// todo: articles
	_, err := database.GetArticles()
	if err != nil {
		log.Fatalln(err)
	}

	//if err := tmplMain.Execute(w, nil); err != nil {
	//	log.Println(err)
	//}
}

// при добавлении статьи, автоматически создавать html, который будет рендериться при переходе по ссылке
