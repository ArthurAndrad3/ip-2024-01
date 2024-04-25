package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var matricula int
	var hora, Vhora float64
	for {
		fmt.Scanf("%d %f %f", &matricula, &hora, &Vhora)
		bufio.NewReader(os.Stdin).ReadByte()
		if matricula == 0 {
			break
		}
		fmt.Printf("%d %.2f\n", matricula, hora*Vhora)

	}
}
