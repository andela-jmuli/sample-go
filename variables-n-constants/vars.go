package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.000000
	b := 3
	ptr := &a

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is type", reflect.TypeOf(b))

	c := int(a) + b

	fmt.Println("\nC has value:", c, "and is of type:", reflect.TypeOf(c))
	fmt.Println("The memory address of variable *a* is ", ptr, "and the value is ", *ptr)
}
