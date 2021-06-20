package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
)

func main() {

	//datos del servidor
	server := "localhost"
	port := "9001"

	// conexion
	conn, error := rpc.DialHTTP("tcp", server+":"+port)
	if error != nil {
		log.Fatal("\n Error en conexion:", error)
	} else {
		log.Printf("\n Cliente RPC en: %s:%s", server, port)
	}

	type Int int

	var reply Int

	fmt.Print("\nIngrese un numero para aumentar o disminuir el contador")
	fmt.Print("\nIngrese la letra [r] para resetear el contador a cero")
	fmt.Print("\nIngrese la letra [p] para ver los procesos en ejecucion\n")

	error = conn.Call("API.Valor", 0, &reply)
	if error != nil {
		log.Fatal("Error:", error)
	}

	var entrada string

	for {
		fmt.Printf("\n Contador: %d\n Entrada: ", reply)
		fmt.Scan(&entrada)

		if i, err := strconv.Atoi(entrada); err != nil {
			//No es un int, revisa si es reset
			if entrada == "r" {
				error = conn.Call("API.Reset", 0, &reply)
				if error != nil {
					log.Fatal("Error:", error)
				}
			} else if entrada == "p" {
				// Es una llamada a ver los procesos
			} else {
				fmt.Println("Ingrese un numero entero o 'r'")
			}
		} else {
			error = conn.Call("API.Aumentar", i, &reply)
			if error != nil {
				log.Fatal("Error:", error)
			}
		}
	}
}
