package demo

import (
	"fmt"
	"github.com/calmh/du"
)

func DuPath(path string) (free int64, err error) {
	u, err := du.Get(path)
	fmt.Println("Path:", path, "FreeBytes:", u.FreeBytes, "TotalBytes:", u.TotalBytes)
	return u.FreeBytes, err
}
