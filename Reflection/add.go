package main

import (
	"reflect"
)

func main() {
	s := reflect.Zero(reflect.TypeOf(reflect.Struct))
	f := reflect.Zero(reflect.TypeOf(reflect.StructField{
		Name:      "id",
		PkgPath:   "",
		Type:      reflect.TypeOf(reflect.Int),
		Tag:       "",
		Offset:    0,
		Index:     []int{},
		Anonymous: false,
	}))

	m := reflect.Zero(reflect.TypeOf(reflect.Method{
		Name:    "add",
		PkgPath: "",
		Type: reflect.TypeOf(reflect.FuncOf(
			[]reflect.Type{reflect.TypeOf(reflect.Int)},
			[]reflect.Type{reflect.TypeOf(reflect.Int)},
			false)),
		Func:  f,
		Index: 0,
	}))

	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	//	t := reflect.ValueOf(&swap).Elem()
	//	m.Set(reflect.MakeFunc(t.Type(), swap))

	//var fptr func(int, int) (int, int)
	fn := reflect.ValueOf(&swap).Elem()
	v := reflect.MakeFunc(fn.Type(), swap)
	fn.Set(v)
	_ = m
	_ = s

}
