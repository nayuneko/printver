package main

import (
	"fmt"
	"github.com/nayuneko/printver"
)

func main() {
	printver.Version()
	fmt.Println("Program start. version:", printver.GetVersion())
}
