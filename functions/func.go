package main

import "fmt"
import "strings"

func main() {
	module := "function basics"
	author := "joseph muli"

	fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
