package demo

import (
	"fmt"
)

func TermColor() {
	fmt.Print("\x1B[31m" + "red" + "\033[0m" + "\n")
	fmt.Print("\x1B[32m" + "green" + "\033[0m" + "\n")
	fmt.Print("\x1B[33m" + "yellow" + "\033[0m" + "\n")
	fmt.Print("\x1B[34m" + "blue" + "\033[0m" + "\n")
	fmt.Print("\x1B[35m" + "mag" + "\033[0m" + "\n")
	fmt.Print("\x1B[36m" + "cyn" + "\033[0m" + "\n")
	fmt.Print("\x1B[37m" + "white" + "\033[0m" + "\n")
}
