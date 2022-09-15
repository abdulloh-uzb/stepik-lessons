package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		if (number < 100 && number >= 10) && (number%8 == 0) {
			sum += number
		}
	}
	fmt.Print(sum)

}
