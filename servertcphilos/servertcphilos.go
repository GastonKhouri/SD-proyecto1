//Servidor TCP con gorutinas (hilos) para manejo de rcp ; puerto 2002

package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
)

//Tipo necesario para hacer los rpc
type Int int

//Variables para conexion
var ip string
var puertotcp string
var puertorpc string

func main() {

	ip := "localhost"
	puertotcp := "2020"

	//Escuchar un request
	ln, err := net.Listen("tcp", ip+":"+puertotcp)
	if err != nil {
		log.Println("Server TCP Hilos: Error en server listen: ", err)
	}

	//Asegura que se cierre cuando termine la conexion
	defer ln.Close()

	log.Println("Entrando a go")
	go escuchar(ln)
	for {
	}

}

func escuchar(ln net.Listener) {
	ip := "localhost"
	puertosalida := "9000"
	var entrada string
	buf := new(bytes.Buffer)
	log.Println("ESCUCHANDO")

	for {
		//Recibir un request
		conne, err := ln.Accept()
		if err != nil {
			log.Println("Server TCP Hilos: Error en server accept", err)
		} else {
			log.Println("Server TCP Hilos: Conexión aceptada")
		}

		//Recibe la entrada como un string
		dec := gob.NewDecoder(conne)
		if err = dec.Decode(&entrada); err != nil {
			log.Println("Server TCP Hilos: Error decodificando: ", err)
		}

		//Comienza conexion con cola

		// Codificar entrada de nuevo
		enc := gob.NewEncoder(buf)
		if err := enc.Encode(entrada); err != nil {
			log.Println("Server TCP Hilos: Error codificando entrada: ", err)
		}

		//Llamar a la cola
		log.Println("Server TCP Hilos: Llamado a cola en:", "localhost:9000")
		conns, err := net.Dial("tcp", ip+":"+puertosalida)
		if err != nil {
			log.Println("Server TCP Hilos: Error en dial: ", err)
		}

		// Escribir a la cola
		buf.WriteTo(conns)
		go escuchar(ln)

		//Lee la conexion y la imprime
		resp, err := bufio.NewReader(conns).ReadString('\n')
		if err != nil {
			fmt.Println("Server TCP Hilos: Error leyendo la conexión de regreso", err)
		}
		log.Println("Server TCP Hilos: Recibido: ", string(resp))

		// Devuelve la respuesta de parte de la cola hacia el cliente
		io.WriteString(conne, fmt.Sprintf(resp))
	}
}
