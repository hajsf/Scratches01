// Golang program to illustrate
// reflect.New() Function

package main

import (
	"fmt"
	"reflect"
)

type Geek struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

// Main function
func main() {
	greeting := "GeeksforGeeks"
	f := Geek{A: 10, B: "Number"}

	gVal := reflect.ValueOf(greeting)

	fmt.Println(gVal.Interface())

	gpVal := reflect.ValueOf(&greeting)
	gpVal.Elem().SetString("Articles")
	fmt.Println(greeting)

	fld := reflect.Zero(reflect.TypeOf(reflect.StructField{
		Name:      "id",
		PkgPath:   "",
		Type:      reflect.TypeOf(reflect.Int),
		Tag:       "",
		Offset:    0,
		Index:     []int{},
		Anonymous: false,
	}))

	fType := reflect.TypeOf(f)
	fVal := reflect.New(fType)
	x := fVal.Elem()
	x.f
	x = reflect.Append(x, fld)
	x.Field(0).SetInt(20)
	x.Field(1).SetString("Number")
	x.Field(2).SetInt(5)
	f2 := fVal.Elem().Interface().(Geek)
	fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)
}
