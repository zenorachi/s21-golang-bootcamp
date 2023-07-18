package credentials

import (
	"bufio"
	"errors"
	"html/template"
	"log"
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
	DBUser        string
}

var (
	AC       AdminCredentials
	TmplAuth = template.Must(template.ParseFiles("./ui/html/authentication.html"))
)

func init() {
	if err := getDataFromCredentials(); err != nil {
		log.Println(err)
	}
}

func getDataFromCredentials() error {
	file, err := os.Open("./internal/credentials/admin_credentials.txt")
	if err != nil {
		return err
	}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		credentials := strings.Split(line, ":")
		switch credentials[0] {
		case "adminLogin":
			AC.AdminLogin = credentials[1]
		case "adminPassword":
			AC.AdminPassword = credentials[1]
		case "dbName":
			AC.DBName = credentials[1]
		case "dbUser":
			AC.DBUser = credentials[1]
		default:
			return errors.New("error in parsing admin credentials")
		}
	}
	return nil
}
