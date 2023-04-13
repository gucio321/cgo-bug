package pkg

import "C"

func Foo() C.int {
	return C.int(5)
}
