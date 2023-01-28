//go:build windows

package main

import (
	"fmt"

	"github.com/Arthur1/windows-api-benchmark/windows"
)

func main() {
	a := windows.GetProcessorArchitectureFromApi()
	fmt.Printf("%d\n", a)
	b := windows.GetProcessorArchitectureFromWmi()
	fmt.Printf("%d\n", b)
	c := windows.GetProcessorArchitectureFromReg()
	fmt.Printf("%s\n", c)
}
