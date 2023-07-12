package main

import (
	"flag"
	"fmt"
)

type Flags struct {
	oldF, newF *string
}

var fl Flags

func init() {
	fl.oldF = flag.String("old", "", "original db")
	fl.newF = flag.String("new", "", "original db")
}

func main() {
	flag.Parse()
	if *fl.oldF == "" || *fl.newF == "" {
		fmt.Println("ERROR:\n\t--old [file name]\n\t--new [file name]")
		return
	}

	compareDumps(*fl.oldF, *fl.newF)
}
