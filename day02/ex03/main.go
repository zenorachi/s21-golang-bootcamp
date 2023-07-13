package main

import (
	"flag"
	"log"
	"os"
)

var pathToDir *string

func init() {
	pathToDir = flag.String("a", "", "path to directory")
	flag.Parse()
	if err := isCorrectFolder(*pathToDir); *pathToDir != "" && err != nil {
		log.Fatal(err)
	}
}

func main() {

}

func isCorrectFolder(dir string) error {
	if info, err := os.Stat(dir); os.IsNotExist(err) || os.IsPermission(err) || !info.IsDir() {
		return err
	}
	return nil
}
