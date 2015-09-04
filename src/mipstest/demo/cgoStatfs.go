package demo

/*
#include <sys/statfs.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//type Statfs_t C.struct_statfs

func CStatf(path string) {
	s := &Statfs_t{}
	s_p := unsafe.Pointer(s)
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.statfs(cpath,
		// convert go pointer to c pointer to struct
		(*C.struct_statfs)(s_p),
	)
	fmt.Println("",
		"Type:", s.Type,
		"Bsize", s.Bsize,
		"Frsize", s.Frsize,
		"Blocks", s.Blocks,
		"Bfree", s.Bfree,
		"Files", s.Files,
		"Ffree", s.Ffree,
		"Bavail", s.Bavail,
		"Namelen", s.Namelen,
		"Flags", s.Flags,
	)
}
