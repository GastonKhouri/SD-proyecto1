// Ejecutable al que llama el servidor TCP por procesos al crear un nuevo proceso

package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

//Tipo necesario para hacer los rpc
type Int int

//Variables para conexion
var ip string
var puertotcp string
var puertorpc string

//Funciones para llamar los RPC
func manejarAumento(n int32, puerto string) {

	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Procedimiento Manejar Contador: Error de conexión: ", err)
	}

	client.Call("API.Aumentar", n, &resp)
	fmt.Printf("Procedimiento Manejar Contador: Hice una llamada para aumentar y devolvió %d \n", resp)
	client.Close()

}

func manejarValor(puerto string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Procedimiento Manejar Contador: Error de conexión: ", err)
	}

	client.Call("API.Valor", 0, &resp)
	fmt.Printf("Procedimiento Manejar Contador: El valor actual es: %d \n", resp)
	client.Close()
}

func manejarReseteo(puerto string) {
	var resp Int

	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

	if err != nil {
		log.Fatal("Procedimiento Manejar Contador: Error de conexión", err)
	}

	client.Call("API.Reset", 0, &resp)
	fmt.Printf("Procedimiento Manejar Contador: Contador reseteado. Valor actual: %d \n", resp)
	client.Close()
}

func main() {

	puertorpc := "9001"
	var num int32

	//Ingresa el argumento a "entrada"
	entrada := os.Args[1]

	//Intenta convertir la entrada a int, si no puede, la entrada es r entonces resetea
	if x, err := strconv.Atoi(entrada); err == nil {
		//No hubo error convirtiendo porque es int
		num = int32(x)
		if num == 0 {
			manejarValor(puertorpc)
		} else {
			manejarAumento(num, puertorpc)
		}
	} else {
		//Hubo error convirtiendo por lo tanto es 'r'

		manejarReseteo(puertorpc)
	}

}
