//Servidor TCP por procesos ; puerto 2002
package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"os/exec"
)

//Tipo necesario para hacer los rpc
type Int int

//Variables para conexion
var ip string
var puertotcp string
var puertorpc string

//Funciones para llamar los RPC
func manejarAumento(conn net.Conn, n int32, puerto string) {

	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Error de conexión: ", err)
	}

	client.Call("API.Aumentar", n, &resp)
	client.Close()

}

func manejarValor(conn net.Conn, puerto string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Error de conexión: ", err)
	}

	client.Call("API.Valor", 0, &resp)
	client.Close()
}

func manejarReseteo(conn net.Conn, puerto string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Error de conexión", err)
	}

	client.Call("API.Reset", 0, &resp)
	client.Close()
}

func main() {

	ip := "localhost"
	puertotcp := "2002"
	var entrada string

	//Escuchar un request
	ln, err := net.Listen("tcp", ip+":"+puertotcp)
	if err != nil {
		log.Println("Server TCP Proc: Error en server listen: ", err)
	} else {
		log.Println("Server TCP Proc: Escuchando en: ", ip+":"+puertotcp)
	}

	//Asegura que se cierre cuando termine la conexion
	defer ln.Close()

	for {

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

		manContOut, err := manContCmd.Output()
		if err != nil {
			log.Println("Server TCP Proc: Error en comando: ", err)
		}

		//Escribir a la conexión la respuesta del proceso ejecutado
		io.WriteString(conn, fmt.Sprintf(string(manContOut)))

	}

}
