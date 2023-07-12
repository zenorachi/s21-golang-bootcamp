package main

import (
	walkfn "day02/ex00/walk-funcs"
	"errors"
	"flag"
	"log"
	"path/filepath"
)

type Flags struct {
	f, d, sl bool
	ext      string
}

var fl Flags

func init() {
	flag.BoolVar(&fl.f, "f", false, "file")
	flag.BoolVar(&fl.d, "d", false, "directory")
	flag.BoolVar(&fl.sl, "sl", false, "symbolic links")
	flag.StringVar(&fl.ext, "ext", "", "extension")
}

func main() {
	flag.Parse()

	filePath, err := checkFlags()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(filePath, walkfn.WalkFunc)
	if err != nil {
		log.Fatal(err)
	}
}

func checkFlags() (string, error) {
	args := flag.Args()
	if len(args) == 0 {
		return "", errors.New("ERROR: missing file path")
	}

	if !fl.f && fl.ext != "" {
		return "", errors.New("ERROR: -ext may only be used with -f")
	}

	if !fl.f && !fl.d && !fl.sl {
		fl.f, fl.d, fl.sl = true, true, true
	}

	return args[0], nil
}
