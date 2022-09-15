package main

import (
	"fmt"
)

func main() {

	var a int // 643198
	fmt.Scan(&a)
	temp1 := (a / 100000) + ((a / 10000) % 10) + ((a % 10000) / 1000)
	b := (a % 1000)
	temp2 := (b / 100) + (b%100)/10 + (b % 10)
	if temp1 == temp2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
