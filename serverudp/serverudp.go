// Servidor UDP ; puerto 2002

package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

//Tipo necesario para hacer los rpc
type Int int

//Variables para conexion
var ip string
var puertoudp string
var puertorpc string

//Funcion para manejar conexiones udp
func handleUDPConnection(conne *net.UDPConn) {

	port := "9000"

	buffer := make([]byte, 1024)
	buf := new(bytes.Buffer)

	n, addr, err := conne.ReadFromUDP(buffer)
	if err != nil {
		log.Println("Server UDP: Error en readfromudp: ", err)
	}

	entrada := bytes.NewBuffer(buffer[:n]).String()

	//Aqui comienza la llamada a la cola

	// Codificar entrada de nuevo
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(entrada); err != nil {
		log.Println("Server UDP: Error codificando entrada: ", err)
	}

	//Llamar a la cola
	log.Println("Server UDP: Llamado a cola en:", ip+":"+port)
	conns, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Println("Server UDP: Error en dial: ", err)
	}

	defer conns.Close()

	// Escribir a la cola
	buf.WriteTo(conns)

	//Lee la conexion y la imprime
	resp, err := bufio.NewReader(conns).ReadString('\n')
	if err != nil {
		fmt.Println("Server UDP: Error leyendo la conexión de regreso", err)
	}
	log.Println("Server UDP: Recibido: ", string(resp))

	//Aqui termina la conexion a la cola

	_, err = conne.WriteToUDP([]byte(resp), addr)
	if err != nil {
		log.Println("Error escribiendo al cliente: ", err)
	}

}

//Funcion para mandar a cola
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
