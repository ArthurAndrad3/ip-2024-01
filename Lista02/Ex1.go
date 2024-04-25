package main

import "fmt"

func main() {
	var matricula, NP, LE, TF, P float64
	for {
		NP, LE = 0, 0
		fmt.Printf("Digite a matrícula do aluno: ")
		fmt.Scan(&matricula)
		for i := 1; i <= 8; i++ {
			var n float64
			fmt.Printf("Digite a nota da prova %d : ", i)
			fmt.Scan(&n)
			NP += n
		}
		NP /= 8
		for i := 1; i <= 5; i++ {
			var n float64
			fmt.Printf("Digite a nota da lista %d: ", i)
			fmt.Scan(&n)
			LE += n
		}
		LE /= 5
		fmt.Printf("Digite a nota do trabalho final: ")
		fmt.Scan(&TF)
		fmt.Printf("Digite a carga horária de presença: ")
		fmt.Scan(&P)
		NF := 0.7*NP + 0.15*LE + 0.15*TF
		if matricula == -1 && NP == -1 && LE == -1 && TF == -1 && P == -1 {
			break
		} else {
			if NF >= 6 && P >= 0.75*128 {
				fmt.Printf("APROVADO\n")
			}
			if NF < 6 && P >= 0.75*128 {
				fmt.Printf("REPROVADO POR NOTA\n")
			}
			if NF >= 6 && P < 0.75*128 {
				fmt.Printf("REPROVADO POR FALTA\n")
			}
			if NF < 6 && P < 0.75*128 {
				fmt.Printf("REPROVADO POR NOTA E FALTA\n")
			}
		}
	}
}
