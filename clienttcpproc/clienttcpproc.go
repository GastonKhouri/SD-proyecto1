//Cliente TCP basado en procesos; puerto 2002

package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {

	go ingresar()
	for {
	}

}

func ingresar() {

	ip := "localhost"
	port := "2002"

	buf := new(bytes.Buffer)
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

	//Encodear la entrada valida
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(entrada); err != nil {
		log.Println("Cliente TCP Proc: Error encodeando entrada: ", err)
	}

	//Llamar al servidor
	log.Println("Cliente TCP Proc: Llamado a:", ip+":"+port)
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Println("Cliente TCP Proc: Error en dial: ", err)
	}

	//Asegura que se cierre la conexión
	defer conn.Close()

	//Escribir buffer a la conexión
	buf.WriteTo(conn)
	go ingresar()

	//Lee la conexion y la imprime
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Cliente TCP: Error leyendo la conexión de regreso", err)
	}
	resp = strings.Replace(resp, "/", "\n", -1)
	log.Println("Cliente TCP Proc: Recibido: ", string(resp))
}
