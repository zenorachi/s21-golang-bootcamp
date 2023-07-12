package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	reader "s21compare/ex00/db-reader"
)

func main() {
	var content reader.DataBase
	var err error
	fileName := flag.String("f", "", "filename")
	flag.Parse()

	content, err = reader.GetDB(*fileName)
	if err != nil {
		panic(err)
	}
	printDB(content, "json")
}

func printDB(db reader.DataBase, format string) {
	var content []byte
	var err error
	if format == "json" {
		content, err = json.MarshalIndent(db, "", "    ")
	} else if format == "xml" {
		content, err = xml.MarshalIndent(db, "", "    ")
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n", content)
}
