//Cliente TCP basado en procesos; puerto 2002

package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fcont"
	"fmt"
	"log"
	"net"
)

func main() {

	ip := "localhost"
	port := "2002"

	// var num int32
	buf := new(bytes.Buffer)

	for {

		//Pedir entrada hasta que se ingrese un entero o 'r'
		entrada := fcont.ChequearEntrada()

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

		//Lee la conexion y la imprime
		resp, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Cliente TCP: Error leyendo la conexión de regreso", err)
		}
		log.Println("Cliente TCP Proc: Recibido: ", string(resp))
	}
}
