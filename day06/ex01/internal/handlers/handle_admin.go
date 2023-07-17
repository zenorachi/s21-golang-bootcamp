package handlers

import (
	"bufio"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type AdminContent struct {
	Login    string
	Password string
	Success  bool
}

type AdminCredentials struct {
	AdminLogin    string
	AdminPassword string
	DBName        string
	DBTable       string
}

var (
	ac       AdminCredentials
	tmplAuth = template.Must(template.ParseFiles("./ui/html/admin.html"))
)

func init() {
	if err := getDataFromCredentials(); err != nil {
		log.Println(err)
	}
}

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	data := AdminContent{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}

	if validateAdmin(data.Login, data.Password) {
		data.Success = true
		log.Println("success")
	} else {
		log.Println("not success")
	}
	tmplAuth.Execute(w, data)
}

func validateAdmin(login, password string) bool {
	if login == ac.AdminLogin && password == ac.AdminPassword {
		return true
	}
	return false
}

func getDataFromCredentials() error {
	file, err := os.Open("./internal/admin_credentials.txt")
	if err != nil {
		return err
	}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		credentials := strings.Split(line, ":")
		switch credentials[0] {
		case "adminLogin":
			ac.AdminLogin = credentials[1]
		case "adminPassword":
			ac.AdminPassword = credentials[1]
		case "dbName":
			ac.DBName = credentials[1]
		case "dbTable":
			ac.DBTable = credentials[1]
		default:
			return errors.New("error in parsing admin credentials")
		}
	}
	return nil
}
