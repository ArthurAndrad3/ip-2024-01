package main

import "fmt"

func main() {
	var n, i, j, sp, si int
	for {
		print("Digite um número: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		} else {
			if n%2 == 0 {
				sp = sp + n
				i = i + 1
			} else {
				si = si + n
				j = j + 1
			}
		}
	}
	var mp, mi float64
	mp = float64(sp) / float64(i)
	mi = float64(si) / float64(j)
	fmt.Printf("MEDIA PAR: %.2f\n MEDIA IMPAR: %.2f.\n", mp, mi)
}
