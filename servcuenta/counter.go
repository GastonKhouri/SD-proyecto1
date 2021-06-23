// Contador global, puerto 9001

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"sync"
)

type API int

type Int int

type Cuenta struct {
	mutex sync.RWMutex
	valor Int
}

var cuenta = Cuenta{valor: 0}

func main() {

	go ServidorRCP()

	fmt.Print("\nConsola local\n")
	fmt.Print("\nIngrese un numero para aumentar o disminuir el contador")
	fmt.Print("\nIngrese la letra [r] para resetear el contador a cero\n")

	var entrada string

	fmt.Printf("\n Contador Actual: %d\n", cuenta.Obtener())

	for {
		fmt.Scan(&entrada)

		if i, err := strconv.Atoi(entrada); err != nil {
			//No es un int, revisa si es reset
			if entrada == "r" {
				cuenta.Reset()
				fmt.Printf("Hice una llamada para resetear y devolvio: %d\n", cuenta.Obtener())
			} else if entrada == "p" {

				//Lista de procesos para filtrar

			} else {
				fmt.Println("Ingrese un numero entero o 'r'")
			}
		} else {
			cuenta.Aumentar(Int(i))
			fmt.Printf("Hice una llamada para aumentar y devolvio: %d\n", cuenta.Obtener())

		}
	}
}

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

	rpc.HandleHTTP()

	//Escuchar por conexion
	ls, err := net.Listen("tcp", ip+":"+puertorpc)

	if err != nil {
		log.Fatal("Counter: Error de listener. ", err)
	}

	//Imprimir que se esta sirviendo y en cual puerto
	log.Printf("Counter: sirviendo RPC en:  %s:%s", ip, puertorpc)
	err = http.Serve(ls, nil)
	if err != nil {
		log.Fatal("Counter: error sirviendo: ", err)
	}

	// se asegura de cerrar el servidor al finalizar
	defer ls.Close()
}

// Procedimientos para contador global
// funcion que aumenta o disminuye el contador global dado un numero
func (c *Cuenta) Aumentar(num Int) {
	c.mutex.Lock()
	c.valor = c.valor + num
	c.mutex.Unlock()
}

// funcion que establece el valor del contador global en cero
func (c *Cuenta) Reset() {
	c.mutex.Lock()
	c.valor = 0
	c.mutex.Unlock()
}

// funcion que devuelve el valor del contador global
func (c *Cuenta) Obtener() Int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.valor
}

// Procedimientos remotos
// Devuelve el valor del contador
func (a *API) Valor(n Int, resp *Int) error {
	log.Println("Counter: Ejecutando procedure consultar")

	*resp = cuenta.Obtener()

	return nil
}

// Aumenta el contador por n y devuelve el nuevo valor del contador
func (a *API) Aumentar(n Int, resp *Int) error {
	log.Println("Counter: Ejecutando procedure aumentar")
	cuenta.Aumentar(n)
	*resp = cuenta.Obtener()

	return nil
}

//Resetea el contador, no hace nada con n y devuelve el nuevo valor
func (a *API) Reset(n Int, resp *Int) error {
	log.Println("Counter: Ejecutando procedure resetar")
	cuenta.Reset()
	*resp = cuenta.Obtener()

	return nil
}
