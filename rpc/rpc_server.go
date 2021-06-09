package rpc_server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// objeto que aloja los metodos
type Op int

// metodos del objeto
func (t *Op) Count(number int, reply *int) error {
	*reply = number + 2
	return nil
}

// funcion que inicia el server
func Server() {
	//datos del server
	port := "1234"
	op := new(Op)

	// registro del objeto para el rpc
	e := rpc.Register(op)
	// por si sale mal el registro
	if e != nil {
		log.Fatal("error registering ")
	}

	// a partir de aqui no se que hice XD
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")

	if e != nil {
		log.Fatal("listener error:", e)
	} else {
		log.Printf("Server RPC exitoso hermano en %s", port)
	}

	e = http.Serve(l, nil)

	if e != nil {
		log.Fatal("serving error:", e)
	}
}
