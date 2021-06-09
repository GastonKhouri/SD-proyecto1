package rpc_client

import (
	"fmt"
	"log"
	"net/rpc"
)

func Client() {

	//datos del servidor
	server := "localhost"
	port := "1234"

	// incializacion del cliente, seria cuestion de
	// añadir que corra infitamente y pedir entradas
	c, e := rpc.DialHTTP("tcp", server+":"+port)
	if e != nil {
		log.Fatal("dialing:", e)
	} else {
		log.Printf("Cliente exitoso hermano en %s:%s", server, port)
	}

	// variables fijas de prueba
	var A = 4
	var reply int

	//llamada a procedimeinto remoto
	e = c.Call("Op.Count", A, &reply)

	fmt.Println(reply)

	// esto es por si hay un error en la llamada al procedimiento
	if e != nil {
		log.Fatal("Error:", e)
	}
}
