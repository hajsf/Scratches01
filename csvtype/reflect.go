package main

import (
	"reflect"
)

func CreateStruct() reflect.Value {
	f := []reflect.StructField{}
	fields := []string{"Myfield", "Area", "Size"}
	types := []interface{}{25, 5, 10.3}

	for i, v := range fields {
		x := reflect.StructField{
			Name: reflect.ValueOf(v).Interface().(string),
			Type: reflect.TypeOf(types[i]),
		}

		f = append(f, x)
	}

	t := reflect.StructOf(f)

	e := reflect.New(t).Elem()

	return e
}
