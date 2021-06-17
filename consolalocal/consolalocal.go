//Servidor TCP por procesos ; puerto 2002
package main

import (
	"fcont"
	"fmt"
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

	var entrada string

	for {

		entrada = fcont.ChequearEntrada()

		//Aqui debe comenzar el fork
		manContCmd := exec.Command("manejarContador", entrada)
		log.Println(fmt.Sprintf("Consola Local: Proceso lanzado, entrada: %s", entrada))

		manContOut, err := manContCmd.Output()
		if err != nil {
			log.Println("Consola Local: Error en comando: ", err)
		}

		//Imprimir el resulatdo
		log.Println(string(manContOut))

	}

}
