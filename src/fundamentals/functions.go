package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "rajendran"
	location := "chennai"

	fmt.Println(converter(name, location))
	bestFinish := bestLeagueFinishes(13, 15, 18, 19, 14, 10)
	fmt.Println(bestFinish)
}

func converter(name, location string) (s1, s2 string) {
	name = strings.Title(name)
	location = strings.ToUpper(location)

	return name, location
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
