package main

import (
	"flag"
	"log"
	reader "s21compare/ex00/db-reader"
	comparer "s21compare/ex01/db-comparer"
)

func main() {
	var oldF, newF *string
	var oldDB, newDB reader.DataBase
	var err error

	oldF = flag.String("old", "", "original db")
	newF = flag.String("new", "", "original db")
	flag.Parse()

	oldDB, err = reader.GetDB(*oldF)
	newDB, err = reader.GetDB(*newF)
	if err != nil {
		log.Fatal(err)
		return
	}

	comparer.CompareDB(oldDB, newDB)
}

// Перевести все поля в мапы
