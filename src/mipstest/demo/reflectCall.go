package demo

import (
	"fmt"
	"reflect"
)

func hello() {
	fmt.Println("hello world")
}

func ReflectCall() {
	h1 := hello
	fv := reflect.ValueOf(h1)
	fv.Call(nil)
}
