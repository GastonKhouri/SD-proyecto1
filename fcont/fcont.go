// Funciones utiles

package fcont

import (
	"fmt"
	"strconv"
)

// Chequea que la entrada es valida ; solo loopea hasta que se ingrese algo valido, se utiliza en todos los clientes
// Devuelve un string
func ChequearEntrada() string {

	var entrada string

	//Pedir entrada hasta que se ingrese un entero o 'r'
	for {
		fmt.Scanln(&entrada)

		if _, err := strconv.Atoi(entrada); err != nil {
			//No es un int, revisa si es reset
			if entrada == "r" {
				// Es un reset
				break
			} else if entrada == "p" {
				// Es una llamada a ver los procesos
				break
			} else {
				fmt.Println("Ingrese un numero entero o 'r'")
			}
		} else {
			// Es un int
			break
		}

	}
	return entrada
}
