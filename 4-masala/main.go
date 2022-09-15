package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	array := []int{}
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		array = append(array, a)
	}

	for idx, item := range array {
		if idx%2 == 0 {
			fmt.Printf("%d ", item)
		}
	}
}
