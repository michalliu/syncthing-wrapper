package demo

import (
	"fmt"
	"reflect"
)

func ReflectMakeFunc() {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}
	makeSwap := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		fn.Set(reflect.MakeFunc(fn.Type(), swap))
	}
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
}
