package main

import (
	"fmt"
)

func main() {
	// map[keyType]valueType
	leagueTitles := make(map[string]int)

	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4
	leagueTitles["DeleteThis"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("\nLeague title: %v\nRecent head to heads: %v\n", leagueTitles, recentHead2Head)

	// Edit and Add new
	leagueTitles["DeleteThis"] = 10
	leagueTitles["newItem"] = 6

	// it's unordered list
	for key, value := range leagueTitles {
		fmt.Printf("\nKey is: %v Value is: %v\n", key, value)
	}
	// How to delete from Map
	delete(leagueTitles, "DeleteThis")
	fmt.Println(leagueTitles)

}
