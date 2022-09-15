package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	var i int = 1
	for i <= n {
		if i%c == 0 && i%d != 0 {
			fmt.Print(i)
			break
		}
		i++
	}

}
