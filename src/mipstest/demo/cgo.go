package demo

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char *s) {
	printf("%s", s);
}
*/
import "C"
import "unsafe"

func CPrint() {
	cs := C.CString("hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
