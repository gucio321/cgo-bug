package main

import "C"
import (
	"fmt"

	"example.com/pkg"
)

func Bar() C.int {
	return pkg.Foo()
}

func main() {
	fmt.Println(Bar())
}
