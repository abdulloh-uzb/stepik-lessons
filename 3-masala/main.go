package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	big_number := array[0]

	for idx := range array {
		if big_number < array[idx] {
			big_number = array[idx]
		}

	}
	fmt.Println(big_number)
}
