package main

import (
	"fmt"
	"os"
	"reflect"
)

// These are global variable to the main package
// var (
// 	name   = "Raj"
// 	course = "Golang"
// 	module = 3.6
// )

func main() {
	// shortform of initializing variables without type
	name := os.Getenv("USERNAME")
	module := 3.6
	course := "Golang"
	ptr := &module
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Memeory address of the *module* variable", ptr, "and value of the *module* variable is", *ptr)

	// passing by value
	changeCourse(course)
	fmt.Println("Your watching the course", course)

	// passing by reference
	changeCourseByReference(&course)
	fmt.Println("Your watching the course", course)

	// const
	const piValue = 3.175
	fmt.Println("Const value", piValue)
}

func changeCourse(course string) string {
	course = "Something changed"
	fmt.Println("I'm trying to change the course name to", course)
	return course
}

func changeCourseByReference(course *string) string {
	*course = "Something changed"
	fmt.Println("I'm trying to change the course name to", *course)
	return *course
}
