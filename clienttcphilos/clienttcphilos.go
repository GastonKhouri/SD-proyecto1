//Cliente TCP basado en hilos; puerto 2020

package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {

	ip := "localhost"
	port := "2020"

	//Declarar y pedir el numero que se enviara
	var entrada string
	// var num int32
	buf := new(bytes.Buffer)

	for {

		//Pedir entrada hasta que se ingrese un entero o 'r'
		for {
			fmt.Scanln(&entrada)

			if _, err := strconv.Atoi(entrada); err != nil {
				//No es un int, revisa si es reset
				if entrada == "r" {
					//Es un reset
					break
				} else {
					fmt.Println("Ingrese un numero entero o 'r'")
				}
			} else {
				// Es un int
				break
			}
		}

		//Encodear la entrada valida
		enc := gob.NewEncoder(buf)
		if err := enc.Encode(entrada); err != nil {
			log.Println("Cliente TCP Hilos: Error encodeando entrada: ", err)
		}

		//Llamar al servidor
		log.Println("Cliente TCP Hilos: Llamado a:", ip+":"+port)
		conn, err := net.Dial("tcp", ip+":"+port)
		if err != nil {
			log.Println("Cliente TCP Hilos: Error en dial: ", err)
		}

		//Asegura que se cierre la conexión
		defer conn.Close()

		//Escribir buffer a la conexión
		buf.WriteTo(conn)

		//Lee la conexion y la imprime
		resp, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Cliente TCP Hilos: Error leyendo la conexión de regreso", err)
		}
		log.Println("Cliente TCP Hilos: Recibido: ", string(resp))
	}
}
