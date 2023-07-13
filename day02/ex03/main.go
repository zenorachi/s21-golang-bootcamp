package main

import (
	"day02/ex03/archiver"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

var pathToDir *string

func init() {
	pathToDir = flag.String("a", ".", "path to directory")
	flag.Parse()
	if err := isCorrectFolder(*pathToDir); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	files, err := getFiles()
	if err != nil {
		log.Fatalln(err)
	}

	wg := new(sync.WaitGroup)
	for _, file := range files {
		wg.Add(1)
		go archiver.ArchiveFile(*pathToDir, file, wg)
	}
	wg.Wait()
}

func isCorrectFolder(dir string) error {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) || os.IsPermission(err) {
		return err
	}
	if !info.IsDir() {
		return errors.New(fmt.Sprintf("stat %s: path must point to directory", dir))
	}
	return nil
}

func getFiles() ([]string, error) {
	var files []string
	if *pathToDir == "." {
		files = os.Args[1:]
	} else {
		files = os.Args[3:]
	}

	if len(files) == 0 {
		return nil, errors.New("error: need to specify 1 or more files")
	}
	return files, nil
}
