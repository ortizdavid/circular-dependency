package module1

import (
	"circular-deps/module2"
	"fmt"
)

func PrintMessage() {
	fmt.Println("Message from Module 1")
}

func UseDependency() {
	module2.PrintMessage()
}