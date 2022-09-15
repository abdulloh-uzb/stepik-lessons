package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	fmt.Printf("It is %d hours %d minutes.", d/30, d*2%60)
}
