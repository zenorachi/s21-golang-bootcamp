package handlers

import (
	"net/http"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/article1.html" {
		http.ServeFile(w, r, "./ui/html/article1.html")
	}
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./ui/html/articles.html")
	}

}
