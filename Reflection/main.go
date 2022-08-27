package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type NativeCommandEngine struct{}

func (nse NativeCommandEngine) Method1() {
	fmt.Println("INFO: Method1 executed!")
}
func (nse NativeCommandEngine) Exit() {
	fmt.Println("INFO: Exit method executed!")
	os.Exit(0)
}
func (nse NativeCommandEngine) callMethodByName(methodName string) {
	method := reflect.ValueOf(nse).MethodByName(methodName)
	if !method.IsValid() {
		fmt.Println("ERROR: \"" + methodName + "\" is not implemented")
		return
	}
	method.Call(nil)
}
func (nse NativeCommandEngine) ShowCommands() {
	val := reflect.TypeOf(nse)
	for i := 0; i < val.NumMethod(); i++ {
		fmt.Println(val.Method(i).Name)
	}
}

/*
func (nse NativeCommandEngine) AddCommands(fn func()) {
	val := reflect.TypeOf(nse)
	val.Key().Method(val.NumMethod()) = fn
}
*/
var Quit = func() { // (nse NativeCommandEngine)
	fmt.Println("INFO: Quit method executed!")
	os.Exit(0)
}

func main() {
	nse := NativeCommandEngine{}
	fmt.Println("A simple Shell v1.0.0")
	fmt.Println("Supported commands:")
	nse.ShowCommands()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for scanner.Scan() {
		nse.callMethodByName(scanner.Text())
		fmt.Print("$ ")
	}
}
