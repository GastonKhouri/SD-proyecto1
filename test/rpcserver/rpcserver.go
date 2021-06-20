package rpcserver

import (
	"cont"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

func ServidorRCP() {
	// variables de servidor
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

// Devuelve el valor del contador
func (a *API) Valor(n int, resp *int) error {
	log.Println("Counter: Ejecutando procedure consultar")

	*resp = cont.Contador.Obtener()

	return nil
}

// Aumenta el contador por n y devuelve el nuevo valor del contador
func (a *API) Aumentar(n int, resp *int) error {
	log.Println("Counter: Ejecutando procedure aumentar")
	cont.Contador.Aumentar(n)
	*resp = cont.Contador.Obtener()

	return nil
}

//Resetea el contador, no hace nada con n y devuelve el nuevo valor
func (a *API) Reset(n int, resp *int) error {
	log.Println("Counter: Ejecutando procedure resetar")
	cont.Contador.Reset()
	*resp = cont.Contador.Obtener()

	return nil
}
