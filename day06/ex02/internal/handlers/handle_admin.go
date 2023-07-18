package handlers

import (
	"day06/ex02/internal/credentials"
	"day06/ex02/internal/database"
	"day06/ex02/internal/fs"
	"log"
	"net/http"
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
			if err := fs.CreateArticle(articleLink); err != nil {
				log.Println(err)
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
