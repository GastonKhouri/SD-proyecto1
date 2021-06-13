// Contador global, puerto 9001

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Int int
type API int

var cont Int
var ip string
var puertorpc string

func main() {

	ip := "localhost"
	puertorpc := "9001"

	//Registrar el api
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Counter: error registering API. ", err)
	}

	//?
	rpc.HandleHTTP()

	//Escuchar por conexion
	ls, err := net.Listen("tcp", ip+":"+puertorpc)

	if err != nil {
		log.Fatal("Counter: Error de listener. ", err)
	}

	//Imprimir que se esta sirviendo y en cual puerto
	log.Printf("Counter: sirviendo RPC en puerto %s", puertorpc)
	err = http.Serve(ls, nil)
	if err != nil {
		log.Fatal("Counter: error sirviendo: ", err)
	}

}

func (a *API) Valor(n Int, resp *Int) error {
	log.Println("Counter: Ejecutando procedure consultar")
	*resp = cont

	return nil
}

// Aumenta el contador por n y devuelve el nuevo valor del contador
func (a *API) Aumentar(n Int, resp *Int) error {
	log.Println("Counter: Ejecutando procedure aumentar")
	cont = cont + n
	*resp = cont

	return nil
}

//Resetea el contador, no hace nada con n y devuelve el nuevo valor
func (a *API) Reset(n Int, resp *Int) error {
	log.Println("Counter: Ejecutando procedure resetar")
	cont = 0
	*resp = cont

	return nil
}
