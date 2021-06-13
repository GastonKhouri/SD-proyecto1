// Servidor UDP ; puerto 2002

package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
)

//Tipo necesario para hacer los rpc
type Int int

//Variables para conexion
var ip string
var puertoudp string
var puertorpc string

//Funcion para manejar conexiones udp
func handleUDPConnection(conn *net.UDPConn) {

	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Println("Server UDP: Error en readfromudp: ", err)
	}

	entrada := bytes.NewBuffer(buffer[:n]).String()

	var num int32

	puertorpc := "9001"
	var resp string

	//Intenta convertir la entrada a int, si no puede, la entrada es r entonces resetea
	if x, err := strconv.Atoi(entrada); err == nil {
		//No hubo error convirtiendo porque es int
		num = int32(x)
		if num == 0 {
			manejarValor(conn, puertorpc, &resp)
		} else {
			manejarAumento(conn, num, puertorpc, &resp)
		}
	} else {
		//Hubo error convirtiendo por lo tanto es 'r'
		manejarReseteo(conn, puertorpc, &resp)
	}

	_, err = conn.WriteToUDP([]byte(resp), addr)
	if err != nil {
		log.Println("Error escribiendo al cliente: ", err)
	}

}

//Funciones para llamar los RPC
func manejarAumento(conn net.Conn, n int32, puerto string, respuesta *string) {

	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Error de conexión: ", err)
	}

	client.Call("API.Aumentar", n, &resp)
	log.Println(fmt.Sprintf("Server UDP: Hice una llamada para aumentar y devolvió %d\n", resp))
	*respuesta = fmt.Sprintf("Hice una llamada para aumentar y devolvió %d \n", resp)

	client.Close()

}

func manejarValor(conn net.Conn, puerto string, respuesta *string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Server UDP: Error de conexión: ", err)
	}

	client.Call("API.Valor", 0, &resp)
	log.Println(fmt.Sprintf("Server UDP: El valor actual es: %d\n", resp))
	*respuesta = fmt.Sprintf("El valor actual es: %d \n", resp)
	client.Close()
}

func manejarReseteo(conn net.Conn, puerto string, respuesta *string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Server UDP: Error de conexión", err)
	}

	client.Call("API.Reset", 0, &resp)
	log.Println(fmt.Sprintf("Server UDP: Contador reseteado. Valor actual: %d\n", resp))
	*respuesta = fmt.Sprintf("Contador reseteado. Valor actual: %d \n", resp)
	client.Close()
}

func main() {

	ip := "localhost"
	puerto := "2002"
	service := ip + ":" + puerto

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Println("Server UDP: Error resolviendo: ", err)
	}

	//Escuchar un request
	ln, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		log.Println("Server UDP: Error en server listenudp: ", err)
	}

	//Asegura que se cierre cuando termine la conexion
	defer ln.Close()
	log.Printf("Server UDP: Servidor udp escuchando en %s\n", ln.LocalAddr().String())

	for {

		handleUDPConnection(ln)

	}
}
