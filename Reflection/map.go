// Golang program to illustrate
// reflect.MapOf() Function

package main

import (
	"fmt"
	"reflect"
)

// Main function
func main() {
	//	var m map[string]int
	ta := reflect.ValueOf("age")
	tc := reflect.ValueOf(2)

	//use of MapOf method
	mapType := reflect.MapOf(reflect.TypeOf(reflect.String), reflect.TypeOf(reflect.Int))
	//mapType := reflect.MapOf(reflect.TypeOf("string"), reflect.TypeOf(123))

	mapValue := reflect.MakeMapWithSize(mapType, 0)

	mapValue.SetMapIndex(ta, tc)
	mapValue.SetMapIndex(reflect.ValueOf("hieght"), reflect.ValueOf(10))

	fmt.Println(mapValue)

	keys := mapValue.MapKeys()
	for _, k := range keys {
		c_key := k.Convert(mapValue.Type().Key())
		c_value := mapValue.MapIndex(c_key)
		fmt.Println("key :", c_key, " value:", c_value)
	}
}
