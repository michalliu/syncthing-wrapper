package demo

import (
	"fmt"
	"github.com/pivotal-golang/bytefmt"
	"runtime"
	"syscall"
)

type MemStatus struct {
	All  uint64
	Used uint64
	Free uint64
	Self uint64
}

func MemStat() MemStatus {
	// self mem stat
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := MemStatus{}
	mem.Self = memStat.Alloc
	// system mem stat
	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err == nil {
		mem.All = uint64(sysInfo.Totalram)
		mem.Free = uint64(sysInfo.Freeram)
		mem.Used = mem.All - mem.Free
	}
	fmt.Println("",
		"Self Alloc:", mem.Self, bytefmt.ByteSize(mem.Self), "\n",
		"All:", mem.All, bytefmt.ByteSize(mem.All), "\n",
		"Used:", mem.Used, bytefmt.ByteSize(mem.Used), "\n",
		"Free:", mem.Free, bytefmt.ByteSize(mem.Free),
	)
	return mem
}
