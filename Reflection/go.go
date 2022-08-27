package main

import (
	"reflect"
)

func main() {
	s := reflect.Zero(reflect.TypeOf(reflect.Struct)) //   .New(typ).Elem()
	_ = s
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
	_ = m
	//s.Set(m)

	// f.Elem().SetInt(4)
	// m.Elem().FieldByNameFunc()
	//	reflect.MakeFunc()

	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	/*	var fptr func(int, int) (int, int)
		fn := reflect.ValueOf(&fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), swap)
		fn.Set(v) */
	/*	makeSwap := func(fptr any) { //func() {
			//fn := reflect.ValueOf(reflect.Pointer).Elem()
			fn := reflect.ValueOf(fptr).Elem()
			// Make a function of the right type.
			v := reflect.MakeFunc(fn.Type(), swap)

			// Assign it to the value fn represents.
			fn.Set(v)
		}
		/*
			// Make and call a swap function for ints.
			var intSwap func(int, int) (int, int)
			makeSwap(&intSwap)
			fmt.Println(intSwap(0, 1))

			// Make and call a swap function for float64s.
			var floatSwap func(float64, float64) (float64, float64)
			makeSwap(&floatSwap)
			fmt.Println(floatSwap(2.72, 3.14))
	*/

}
