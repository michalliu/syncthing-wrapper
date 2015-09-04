package demo

import (
	"fmt"
	"runtime"
	"time"
)

func GetCPUCount() {
	print("GOOS:", runtime.GOOS, "\n")
	print("GOARCH:", runtime.GOARCH, "\n")
	print("NumCPU:", fmt.Sprintf("%d", runtime.NumCPU()), "\n")
}

func perfTest(c chan bool, n int) {
	x := 0
	for i := 0; i < 100000000; i++ {
		x += i
	}
	println(n, " ", x)
	if n == 9 {
		c <- true
	}
}

func PerfTest() {
	GetCPUCount()
	//useCore := 1
	//print("USE CORE:", fmt.Sprintf("%d", useCore), "\n")
	//runtime.GOMAXPROCS(useCore)
	println("start")
	t1 := time.Now()
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go perfTest(c, i)
	}
	<-c
	println("done, ", time.Now().Sub(t1).String())
}
