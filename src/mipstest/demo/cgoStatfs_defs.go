// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs src/mipstest/demo/cgoStatfs.go

package demo

type Statfs_t struct {
	Type	int32
	Bsize	int32
	Frsize	int32
	Blocks	uint32
	Bfree	uint32
	Files	uint32
	Ffree	uint32
	Bavail	uint32
	Fsid	_Ctype_struct___0
	Namelen	int32
	Flags	int32
	Spare	[5]int32
}
