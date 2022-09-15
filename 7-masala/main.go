package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	// fmt.Println(string(a[0]), b)
	for _, item := range a {
		for _, item2 := range b {
			if string(item) == string(item2) {
				fmt.Print(string(item), " ")
			}
		}
	}
}
