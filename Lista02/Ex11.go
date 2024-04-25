package main

import "fmt"

func main() {
	var n, e, erro, r float64
	r = 1
	erro = n - r*r
	if erro < 0 {
		erro = -erro
	}
	fmt.Scan(&n)
	fmt.Scan(&e)
	for erro > e {
		r = (r + n/r) / 2
		erro = n - r*r
		if erro < 0 {
			erro = -erro
		}
		fmt.Printf("r = %.9f, erro: %.9f\n", r, erro)
	}
}
