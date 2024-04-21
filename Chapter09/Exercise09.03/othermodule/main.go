package main

import (
	"fmt"

	"github.com/ibiscum/Go-Programming-From-Beginner-to-Professional-2nd/Chapter09/Exercise09.03/printer"
)

func main() {
	msg := printer.PrintNewUUID()
	fmt.Println(msg)
}
