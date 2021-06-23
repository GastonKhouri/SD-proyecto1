//Servidor TCP por procesos ; puerto 2002
package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

//Tipo necesario para hacer los rpc
type Int int

//Variables para conexion
var ip string
var puertotcp string
var puertorpc string

func main() {

	ip := "localhost"
	puertotcp := "2002"

	//Escuchar un request
	ln, err := net.Listen("tcp", ip+":"+puertotcp)
	if err != nil {
		log.Println("Server TCP Proc: Error en server listen: ", err)
	} else {
		log.Println("Server TCP Proc: Escuchando en: ", ip+":"+puertotcp)
	}

	//Asegura que se cierre cuando termine la conexion
	defer ln.Close()

	go recibir(ln)
	for {
	}

}

func recibir(ln net.Listener) {

	var entrada string

	//Recibir un request
	conn, err := ln.Accept()
	if err != nil {
		log.Println("Server TCP Proc: Error en server accept", err)
	} else {
		log.Println("Server TCP Proc: Conexión aceptada")
	}

	//Recibe la entrada como un string
	dec := gob.NewDecoder(conn)
	if err := dec.Decode(&entrada); err != nil {
		log.Println("Server TCP Proc: Error decodificando: ", err)
	}

	//Aqui debe comenzar el fork
	manContCmd := exec.Command("manejarContador", entrada)
	log.Println(fmt.Sprintf("Server TCP Proc: Proceso lanzado, entrada: %s", entrada))

	go recibir(ln)

	manContOut, err := manContCmd.Output()
	if err != nil {
		log.Println("Server TCP Proc: Error en comando: ", err)
	} //Funciona el output puesto en manContOut

	//Escribir a la conexión la respuesta del proceso ejecutado
	io.WriteString(conn, fmt.Sprintf(string(manContOut)))
}
