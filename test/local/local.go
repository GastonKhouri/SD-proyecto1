package main

import (
	"cont"
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("\nConsola local\n")
	fmt.Print("\nIngrese un numero para aumentar o disminuir el contador")
	fmt.Print("\nIngrese la letra [r] para resetear el contador a cero")
	fmt.Print("\nIngrese la letra [p] para ver los procesos en ejecucion\n")

	var entrada string

	for {
		fmt.Printf("\n Contador: %d\n Entrada: ", cont.Contador.Obtener())
		fmt.Scan(&entrada)

		if i, err := strconv.Atoi(entrada); err != nil {
			//No es un int, revisa si es reset
			if entrada == "r" {
				cont.Contador.Reset()
			} else if entrada == "p" {
				// Es una llamada a ver los procesos
			} else {
				fmt.Println("Ingrese un numero entero o 'r'")
			}
		} else {
			cont.Contador.Aumentar(i)
		}
	}
}
