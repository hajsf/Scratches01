package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Book struct {
		name   string
		author string
	}
	sampleBook := Book{"Reflection in Go", "John"}
	fmt.Println(reflect.ValueOf(sampleBook).Field(1)) // John

	fields := reflect.ValueOf(sampleBook).NumField()
	fmt.Println(fields)
}
