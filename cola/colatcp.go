package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"time"
)

type solicitud struct {
	num  string
	conn net.Conn
}

func aceptarRequest(p chan (bool), canalConexion chan (net.Conn), ln net.Listener) {

	var entrada string
	defer ln.Close()

	for {

		//Recibir un request
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Cola: Error en server accept", err)
		} else {
			log.Println("Cola: Conexión aceptada")
		}

		time.Sleep(time.Second * 2)
		// canalConexion <- conn

		//Decodificar el request para agregarlo a la cola como string
		dec := gob.NewDecoder(conn)
		if err = dec.Decode(&entrada); err != nil {
			log.Println("Cola: Error decodificando: ", err)
		}

		//Se tiene ahora un string en entrada
		//Ahora se intentará agregarlo a la cola
		log.Println("Intentando agregar a cola: ", entrada)

		agregarACola(&solicitud{entrada, conn})
	}
}

func enviarReqAContador(p chan (bool), canalConexion chan (net.Conn), puertorpc string) {

	for {

		if len(cola) != 0 {
			llamarRPC(puertorpc)

		}
	}
}

func llamarRPC(ps string) {
	var num int32

	//Se saca el primer elemento
	entrada, conn := dequeue()

	//Intenta convertir la entrada a int, si no puede, la entrada es r entonces resetea
	if x, err := strconv.Atoi(entrada); err == nil {
		//No hubo error convirtiendo porque es int
		num = int32(x)
		if num == 0 {
			manejarValor(conn, ps)
			log.Println("Se manejo valor en cola")
		} else {
			manejarAumento(conn, num, ps)
			log.Println("Se manejo aumento en cola")
		}
	} else {
		//Hubo error convirtiendo por lo tanto es 'r'
		manejarReseteo(conn, ps)
		log.Println("Se manejo reseteo en cola")
	}
}

func agregarACola(i *solicitud) {
	cola = append(cola, *i)
	log.Println("Cola: Agregado a cola: ", *i)

}

func dequeue() (string, net.Conn) {
	element := cola[0]
	cola = cola[1:]
	return element.num, element.conn
}

//Tipo necesario para hacer los rpc
type Int int
type API int

var cola []solicitud
var puertoentrada string
var puertosalida string
var ip string
var permiso chan (bool) //Canal permiso
var cc chan (net.Conn)  //Canal conexion

func main() {

	puertoentrada = "9000"
	puertosalida = "9001"
	ip = "localhost"
	// buf := new(bytes.Buffer)

	//Escuchar un request
	ln, err := net.Listen("tcp", ip+":"+puertoentrada)
	if err != nil {
		log.Println("Cola: Error en server listen: ", err)
	} else {
		log.Println("Escuchando en puerto ", puertoentrada)
	}

	// Recibir a la cola
	go aceptarRequest(permiso, cc, ln)
	// Mandar de la cola al contador
	enviarReqAContador(permiso, cc, puertosalida)

	log.Println("Final de cola")
	// for {

	// }
}

//Funciones para llamar los RPC
func manejarAumento(conn net.Conn, n int32, puerto string) {

	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Error de conexión: ", err)
	}

	client.Call("API.Aumentar", n, &resp)
	log.Printf(fmt.Sprintf("Cola: Hice una llamada para aumentar y devolvió %d\n", resp))
	io.WriteString(conn, fmt.Sprintf("Hice una llamada para aumentar y devolvió %d \n", resp))
	client.Close()

}

func manejarValor(conn net.Conn, puerto string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Cola: Error de conexión: ", err)
	}

	client.Call("API.Valor", 0, &resp)
	log.Printf(fmt.Sprintf("Cola: El valor actual es: %d\n", resp))
	io.WriteString(conn, fmt.Sprintf("Hice una llamada para valor y devolvió %d \n", resp))
	client.Close()
}

func manejarReseteo(conn net.Conn, puerto string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Cola: Error de conexión", err)
	}

	client.Call("API.Reset", 0, &resp)
	log.Printf(fmt.Sprintf("Cola: Contador reseteado. Valor actual: %d\n", resp))
	io.WriteString(conn, fmt.Sprintf("Hice una llamada para resetear y devolvió %d \n", resp))
	client.Close()
}
