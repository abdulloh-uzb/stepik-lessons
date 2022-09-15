package main

import (
	"fmt"
)

func main() {
	var temp int

	for fmt.Scan(&temp); temp <= 100; fmt.Scan(&temp) {

		if temp >= 10 {
			fmt.Println(temp)
		}

	}
}
