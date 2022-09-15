package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	var year int = 0
	for a < c {

		a += b * a / 100 //sum += p * sum / 100
		year++
	}
	fmt.Println(year)
}
