package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Flags struct {
	l, m, w bool
}

var fl Flags

func init() {
	flag.BoolVar(&fl.l, "l", false, "count lines")
	flag.BoolVar(&fl.m, "m", false, "count characters")
	flag.BoolVar(&fl.w, "w", false, "count words")
}

func main() {
	flag.Parse()

	err := checkFlags()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK")
}

func checkFlags() error {
	if fl.l && fl.m && fl.w {
		return errors.New("ERROR: only one flag can be used")
	}

	if (xor(fl.l, fl.m) && fl.w) || (xor(fl.m, fl.w) && fl.l) {
		return errors.New("ERROR: only one flag can be used")
	}

	return nil
}

func xor(x, y bool) bool {
	return (x || y) && !(x && y)
}
