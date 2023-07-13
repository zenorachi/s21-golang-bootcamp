package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func countThings(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	var l, m, w int

	file, e := os.Open(fileName)
	if e != nil {
		log.Println(e.Error())
		return
	}

	defer func(f *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l++
		m += len(scanner.Bytes())
		line := scanner.Text()
		spaces := strings.Count(line, " ")
		if spaces != 0 {
			w += spaces + 1
		}
	}

	printThing(l, m, w, fileName)
}

func printThing(l, m, w int, f string) {
	if fl.l {
		fmt.Printf("%d\t%s\n", l, f)
	} else if fl.m {
		fmt.Printf("%d\t%s\n", m, f)
	} else if fl.w {
		fmt.Printf("%d\t%s\n", w, f)
	}
}
