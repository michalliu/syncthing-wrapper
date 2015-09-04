package demo

import (
	"fmt"
	"os/exec"
)

func ExecCommand(s string) {
	cmd := exec.Command(s)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
