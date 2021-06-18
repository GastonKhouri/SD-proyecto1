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
	"net/rpc"
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
	puertosalida := "9000"

	//Escuchar un request
	ln, err := net.Listen("tcp", ip+":"+puertotcp)
	if err != nil {
		log.Println("Server TCP Hilos: Error en server listen: ", err)
	}

	//Asegura que se cierre cuando termine la conexion
	defer ln.Close()
	var entrada string
	buf := new(bytes.Buffer)

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
		log.Println("Server TCP Hilos: Llamado a cola en:", ip+":"+puertosalida)
		conns, err := net.Dial("tcp", ip+":"+puertosalida)
		if err != nil {
			log.Println("Server TCP Hilos: Error en dial: ", err)
		}

		defer conns.Close()

		// Escribir a la cola
		buf.WriteTo(conns)

		//Lee la conexion y la imprime
		resp, err := bufio.NewReader(conns).ReadString('\n')
		if err != nil {
			fmt.Println("Server TCP Hilos: Error leyendo la conexión de regreso", err)
		}
		log.Println("Server TCP Hilos: Recibido: ", string(resp))

		io.WriteString(conne, fmt.Sprintf(resp))

		// Termina conexion con cola

		// //Intenta convertir la entrada a int, si no puede, la entrada es r entonces resetea
		// if x, err := strconv.Atoi(entrada); err == nil {
		// 	//No hubo error convirtiendo porque es int
		// 	num = int32(x)
		// 	if num == 0 {
		// 		go manejarValor(conn, puertorpc)
		// 	} else {
		// 		go manejarAumento(conn, num, puertorpc)
		// 	}
		// } else {
		// 	//Hubo error convirtiendo por lo tanto es 'r'
		// 	go manejarReseteo(conn, puertorpc)
		// }
	}
}
