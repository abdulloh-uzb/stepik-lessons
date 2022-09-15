package main

import "fmt"

func main() {

	var i int
	fmt.Scan(&i)
	n1 := i / 100
	n2 := (i / 10) % 10
	n3 := i % 100 % 10

	if n1 == n2 || n2 == n3 || n1 == n3 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}

}
