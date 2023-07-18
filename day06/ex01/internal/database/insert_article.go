package database

import (
	"fmt"
)

const (
	col1 = "articlename"
	col2 = "linktoarticle"
)

func InsertArticle(name, link string) error {
	db, err := runDB()
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO articles (%s, %s) VALUES ($1, $2)", col1, col2)
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, link)
	if err != nil {
		return err
	}
	return nil
}
