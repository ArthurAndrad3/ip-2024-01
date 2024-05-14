package main

import "fmt"

func main() {
	var alunos, minalunos, presenca int
	fmt.Scan(&alunos, &minalunos)
	if alunos < 0 || alunos > 1000 || minalunos < 0 || minalunos > 1000 {
		return
	}
	chegada := make([]int, alunos)
	for i := 0; i < alunos; i++ {
		fmt.Scan(&chegada[i])
		if chegada[i] <= 0 {
			presenca++
		}
	}
	if presenca < minalunos {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
		for i := alunos - 1; i >= 0; i-- {
			if chegada[i] <= 0 {
				fmt.Printf("%d\n", i+1)
			}
		}
	}
}
