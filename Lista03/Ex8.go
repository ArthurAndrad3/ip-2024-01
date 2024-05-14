package main

import "fmt"

func main() {
	for {
		var n int
		fmt.Scan(&n)
		if n < 1 {
			break
		}
		fmt.Printf("%b\n", n)
	}
}
