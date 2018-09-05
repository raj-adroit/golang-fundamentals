package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// if statement; condition
	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\n First course is doing better")
	} else {
		fmt.Println("\n Second doing better")
	}
	// switch statement; condition
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got an even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got an odd number -", tmpNum)
	default:
		fmt.Println("Something wrong in the option")
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
