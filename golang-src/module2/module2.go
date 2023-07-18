package module2

import (
	"circular-deps/module1"
	"fmt"
)

func PrintMessage() {
	fmt.Println("Message from Module 2")
}

func UseDependency() {
	module1.PrintMessage()
}