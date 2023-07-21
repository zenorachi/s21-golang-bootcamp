package main

import (
	"errors"
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
	ctfmt "github.com/daviddengcn/go-colortext/fmt"
	"log"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func main() {
	up := UnknownPlant{
		"rose",
		"oval",
		134,
	}
	ap := AnotherUnknownPlant{
		10,
		"lanceolate",
		15,
	}

	ctfmt.Println(ct.Green, true, "UnknownPlant")
	describePlant(up)
	ctfmt.Println(ct.Magenta, true, "\nAnotherUnknownPlant")
	describePlant(ap)
}

func describePlant(plant interface{}) {
	t := reflect.TypeOf(plant)
	if !isStruct(t) {
		log.Fatalln(errors.New("invalid type for describePlant()"))
	}

	v := reflect.ValueOf(plant)

	var (
		fieldType  reflect.StructField
		fieldValue reflect.Value
	)
	for i := 0; i < t.NumField()-1; i++ {
		fieldType = t.Field(i)
		fieldValue = v.Field(i)
		printfField(fieldType, fieldValue)
		fmt.Println(",")
	}
	fieldType = t.Field(t.NumField() - 1)
	fieldValue = v.Field(t.NumField() - 1)
	printfField(fieldType, fieldValue)
	fmt.Println()
}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func printfField(sf reflect.StructField, v reflect.Value) {
	tag := sf.Tag.Get("color_scheme")
	if tag == "" {
		tag = sf.Tag.Get("unit")
	}

	if tag == "" {
		fmt.Printf("%s:%v", sf.Name, v.Interface())
	} else {
		tmpTag := strings.Split(string(sf.Tag), ":")[0]
		fmt.Printf("%s(%s=%v):%v", sf.Name, tmpTag, tag, v.Interface())
	}
}
