package demo

/*
#include <sys/statfs.h>
#include <stdlib.h>

typedef struct demo_statfs_wrapper
{
	long f_bsize;
	long f_blocks;
	long f_bfree;
	long f_bavail;
} demo_statfs_wrapper;

void demo_get_statfs(char* path, demo_statfs_wrapper* stat_w){
	struct statfs stat;
	statfs(path,&stat);
	stat_w->f_bsize = stat.f_bsize;
	stat_w->f_blocks = stat.f_blocks;
	stat_w->f_bfree = stat.f_bfree;
	stat_w->f_bavail = stat.f_bavail;
}
*/
import "C"
import (
	"unsafe"
)

//type DemoStatfsWrapperGo C.demo_statfs_wrapper

func CStatf2(path string) {
	stat := new(DemoStatfsWrapperGo)
	stat_p := unsafe.Pointer(stat)
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.demo_get_statfs(cpath, (*C.demo_statfs_wrapper)(stat_p))
	print(stat.Bsize, "\n")
}
