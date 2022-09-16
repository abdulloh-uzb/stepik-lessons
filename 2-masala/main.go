package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	a := make([]int, 0, 10)
	for i := 0; i < number; i++ {
		var temp int
		fmt.Scan(&temp)
		a = append(a, temp)
	}
	fmt.Print(a[3])
}
