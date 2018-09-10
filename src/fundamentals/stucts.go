package main

import (
	"fmt"
)

func main() {

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive.Author)
	fmt.Println(DockerDeepDive.Level)
}
