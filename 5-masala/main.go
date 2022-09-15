package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var j int = 0
	for i := 0; i < n; i++ {

		var number int
		fmt.Scan(&number)
		if number > 0 {
			j++
		}
	}
	fmt.Println(j)
}
