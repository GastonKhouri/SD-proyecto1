//Servidor TCP con gorutinas (hilos) para manejo de rcp ; puerto 2002

package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"strconv"
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
		log.Fatal("Server TCP Hilos: Error de conexión: ", err)
	}

	client.Call("API.Aumentar", n, &resp)
	log.Println(fmt.Sprintf("Server TCP Hilos: Hice una llamada para aumentar y devolvió %d", resp))
	io.WriteString(conn, fmt.Sprintf("Hice una llamada para aumentar y devolvió %d \n", resp))
	client.Close()

}

func manejarValor(conn net.Conn, puerto string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Server TCP Hilos: Error de conexión: ", err)
	}

	client.Call("API.Valor", 0, &resp)
	log.Println(fmt.Sprintf("Server TCP Hilos: El valor actual es: %d", resp))
	io.WriteString(conn, fmt.Sprintf("El valor actual es: %d \n", resp))
	client.Close()
}

func manejarReseteo(conn net.Conn, puerto string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Server TCP Hilos: Error de conexión", err)
	}

	client.Call("API.Reset", 0, &resp)
	log.Println(fmt.Sprintf("Server TCP Hilos: Contador reseteado. Valor actual: %d", resp))
	io.WriteString(conn, fmt.Sprintf(" Contador reseteado. Valor actual: %d \n", resp))
	client.Close()
}

func main() {

	ip := "localhost"
	puertotcp := "2020"
	puertorpc := "9001"

	//Escuchar un request
	ln, err := net.Listen("tcp", ip+":"+puertotcp)
	if err != nil {
		log.Println("Server TCP Hilos: Error en server listen: ", err)
	}

	//Asegura que se cierre cuando termine la conexion
	defer ln.Close()
	var num int32
	var entrada string

	for {
		//Recibir un request
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Server TCP Hilos: Error en server accept", err)
		} else {
			log.Println("Server TCP Hilos: Conexión aceptada")
		}

		//Recibe la entrada como un string
		dec := gob.NewDecoder(conn)
		if err = dec.Decode(&entrada); err != nil {
			log.Println("Server TCP Hilos: Error decodificando: ", err)
		}

		//Intenta convertir la entrada a int, si no puede, la entrada es r entonces resetea
		if x, err := strconv.Atoi(entrada); err == nil {
			//No hubo error convirtiendo porque es int
			num = int32(x)
			if num == 0 {
				go manejarValor(conn, puertorpc)
			} else {
				go manejarAumento(conn, num, puertorpc)
			}
		} else {
			//Hubo error convirtiendo por lo tanto es 'r'
			go manejarReseteo(conn, puertorpc)
		}
	}
}
