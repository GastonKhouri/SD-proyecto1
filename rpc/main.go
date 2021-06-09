package main

import (
	"rpc_client"
	"rpc_server"
)

// funcion que inicia el server y el cliente para pruebas
func main() {
	// esto ajuro tiene que ser concurrente en este caso
	// para darle paso al cliente para que se ejecute.
	// por los momentos el cliente no debe ser concurrente
	go rpc_server.Server()
	rpc_client.Client()
}
