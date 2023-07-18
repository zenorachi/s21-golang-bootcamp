package handlers

import (
	"day06/ex01/internal/credentials"
	"day06/ex01/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	data := credentials.AdminContent{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}

	if validateAdmin(data.Login, data.Password) {
		http.ServeFile(w, r, "./ui/html/admin.html")
		articleName := r.FormValue("articleName")
		articleLink := r.FormValue("articleLink")
		if articleName != "" && articleLink != "" && strings.Contains(articleLink, ".md") {
			if isCorrectArticle(articleLink) {
				createArticle(articleLink)
			}
			if err := database.InsertArticle(articleName, articleLink); err != nil {
				log.Println(err)
			}
		}
		log.Println("logged in as admin")
	}

	if err := credentials.TmplAuth.Execute(w, data); err != nil {
		log.Println(err)
	}
}

func validateAdmin(login, password string) bool {
	if login == credentials.AC.AdminLogin && password == credentials.AC.AdminPassword {
		return true
	}
	return false
}

func createArticle(name string) {
	mdFile := fmt.Sprintf("./ui/md/%s", name)
	if _, err := os.Create(mdFile); err != nil {
		log.Println(err)
	}

	htmlName := strings.Split(name, ".md")[0] + ".html"
	htmlFile := fmt.Sprintf("./ui/html/%s", htmlName)
	if _, err := os.Create(htmlFile); err != nil {
		log.Println(err)
	}
}

func isCorrectArticle(link string) bool {
	articles, err := database.GetArticles()
	if err != nil {
		log.Println(err)
	}
	for current, _ := range articles {
		if current == link {
			return false
		}
	}
	return true
}
