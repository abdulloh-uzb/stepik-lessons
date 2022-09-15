package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	if i > 0 {
		fmt.Print("Число положительное")
	} else if i == 0 {
		fmt.Print("Ноль")
	} else {
		fmt.Print("Число отрицательное")
	}
}
