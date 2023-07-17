package handlers

import (
	"day06/ex01/internal/credentials"
	"log"
	"net/http"
)

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	data := credentials.AdminContent{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}

	if validateAdmin(data.Login, data.Password) {
		data.Success = true
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
