package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func compareDumps(fileName1, fileName2 string) {
	data1 := getFileMap(fileName1)
	data2 := getFileMap(fileName2)

	for key, value := range data2 {
		if _, ok := data1[key]; !ok {
			fmt.Println("ADDED:", value)
		}
	}

	for key, value := range data1 {
		if _, ok := data2[key]; !ok {
			fmt.Println("REMOVED:", value)
		}
	}

}

func getFileMap(fileName string) map[string]string {
	var file *os.File
	var err error

	file, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		if errClose := f.Close(); errClose != nil {
			log.Fatal(errClose)
		}
	}(file)

	lines := make(map[string]string)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		lines[line] = line
	}
	return lines
}
