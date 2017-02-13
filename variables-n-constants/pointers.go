package main

import (
	"fmt"
)

func main() {
	name := "Jojo"
	course := "Docker deep dive"

	fmt.Println("\nHi", name, "You're currently watching", course)

	// adding an ampersand as a prefix to the course variable assignes a pointer to
	// the "original" course variable
	changeCourse(&course)
	fmt.Println("You are now watching course", course)
}

func changeCourse(course *string) string {
	// the asterix tells Go that the variable is a pointer to a string.

	*course = "First Look: Native Docker Clustering"
	// the asterix tells Go that we want to assign the value to the location in memory
	// that the course pointer is referencing.
	fmt.Println("Trying to change your course to", *course)

	return *course
}
