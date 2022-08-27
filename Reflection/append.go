// Golang program to illustrate
// reflect.Append() Function

package main

import (
	"fmt"
	"reflect"
)

// Main function
func main() {

	var str []string
	var v reflect.Value = reflect.ValueOf(&str)
	//var m map[string]int // m := make(map[string]int)
	//var v reflect.Value = reflect.ValueOf(&m)
	v = v.Elem()

	// using the function
	v = reflect.Append(v, reflect.ValueOf("a"))
	v = reflect.Append(v, reflect.ValueOf("b"))
	v = reflect.Append(v, reflect.ValueOf("c"), reflect.ValueOf("j, k, l"))

	fmt.Println("Our value is a type of :", v.Kind())

	vSlice := v.Slice(0, v.Len())
	vSliceElems := vSlice.Interface()

	fmt.Println("With the elements of : ", vSliceElems)

}
