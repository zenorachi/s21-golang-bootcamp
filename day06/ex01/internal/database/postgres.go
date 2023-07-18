package database

import (
	"database/sql"
	"day06/ex01/internal/credentials"
	"fmt"
	_ "github.com/lib/pq"
)

func runDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("%s://%s:@localhost:8080/myblog?sslmode=disable", credentials.AC.DBUser, credentials.AC.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetArticles() (map[string]string, error) {
	db, err := runDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make(map[string]string)
	for rows.Next() {
		var id int
		var name, link string
		err := rows.Scan(&id, &name, &link)
		if err != nil {
			return nil, err
		}
		articles[link] = name
	}
	return articles, nil
}
