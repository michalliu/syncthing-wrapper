package demo

import (
	"fmt"
	"runtime"
	"time"
)

type Foo struct {
	name string
	num  int
}

func finalizer(f *Foo) {
	fmt.Println("a finalizer has run for ", f.name, f.num)
}

var holder *Foo
var counter int

func MakeFoo(name string) (a_foo *Foo) {
	a_foo = &Foo{name, counter}
	fmt.Println("foo created ", "name=", name, "num=", counter)
	counter++
	if counter == 1{
		runtime.SetFinalizer(a_foo, finalizer)
	}
	return
}

func Bar() {
	f1 := MakeFoo("one")
	if holder == nil {
		holder = f1
		fmt.Println("holder is pointed to ", (*f1).name, (*f1).num)
	}
	MakeFoo("two")
}

func RuntimeSetFinalizer() {
	fmt.Println("SetFinalizer demo start")
	for i := 0; i < 3; i++ {
		Bar()
		fmt.Println("wait for 1 sec")
		time.Sleep(time.Second)
		fmt.Println("run gc")
		runtime.GC()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("done.")
}
