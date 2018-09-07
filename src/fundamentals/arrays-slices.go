package main

import (
	"fmt"
)

func main() {
	// Declares a slice
	myCourses := make([]string, 5, 10)
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceOfSlice := mySlice[2:5]

	fmt.Printf("Length is: %d\nCapacity is: %d\n", len(myCourses), cap(myCourses))
	fmt.Println(mySlice[4])
	fmt.Println(sliceOfSlice)

}
