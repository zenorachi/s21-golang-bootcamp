package handlers

import (
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(rate.Limit(100), 1)

func HandleDefault(w http.ResponseWriter, r *http.Request) {

	if limiter.Allow() {
		if r.URL.Path == "/article1.html" {
			http.ServeFile(w, r, "./ui/html/article1.html")
		}
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./ui/html/articles.html")
		}
	} else {
		http.Error(w, "ot", http.StatusTooManyRequests)
	}
}
