package main

import "fmt"

func main() {

	var i int
	fmt.Scan(&i)
	var j int = i % 100
	j = j / 10
	fmt.Print(j)

}
