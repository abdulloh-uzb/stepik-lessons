package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a < b {
		sum := 0
		for a <= b {
			sum += a
			a++
		}
		fmt.Print(sum)
	}
}
