package main

import (
	"errors"
	"flag"
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

}

func checkFlags() error {
	const errTxt = "ERROR: only one flag can be used"
	if fl.l && fl.m && fl.w {
		return errors.New(errTxt)
	}

	if (xor(fl.l, fl.m) && fl.w) || (xor(fl.m, fl.w) && fl.l) {
		return errors.New(errTxt)
	}

	return nil
}

func xor(x, y bool) bool {
	return (x || y) && !(x && y)
}
