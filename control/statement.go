package main

import (
	"fmt"
	"strings"
)

func main() {
	a, b := 5, 5
	if a == b { // seems like this must be in the same line
		fmt.Println("they match")
	} else { // still
		fmt.Println("not same")
	}
	fmt.Println(strings.Index("chicken", "ken"))
}
