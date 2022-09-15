package main

import (
	"fmt"
)

func main() {

	var i int
	fmt.Scan(&i)
	var a int = 0
	k := i
	for k > 0 {
		k /= 10
		a++
	}
	result := 1
	for j := 1; j < a; j++ {
		result *= 10
	}
	fmt.Print(i / result)

}
