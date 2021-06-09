package main

import (
	"clearscreen"
	"fmt"
)

//Funcion para modificar contador
func modificarContador(n int) {
	contador = contador + n
	agregarALista("127.0.0.1", n)
}

//Leer modificaciones
func LeerModificaciones(edicion chan bool) {
	var n int
	for {
		n = 0
		fmt.Scanln(&n)
		modificarContador(n)
		edicion <- true
	}
}

//Imprimir modificaciones
func impresion(edicion chan bool) {
	var bandera bool
	for {
		//Cada vez que se ingrese una nueva edición se recibe un true, se reimprime y se regresa a false
		bandera = <-edicion
		if bandera == true {
			clearscreen.CallClear()
			imprimirEditores()
			fmt.Printf("Contador: %v \n", contador)
			bandera = false
		}
	}
}

var contador int = 0

func main() {
	//Se declara el canal que va a llevar la información de cuando se realice una edición
	edicion := make(chan bool)

	//Estos dos procesos corren concurrentemente y se comunican por el canal edicion
	go LeerModificaciones(edicion)
	go impresion(edicion)

	//Repite infinitamente
	for {

	}

}
