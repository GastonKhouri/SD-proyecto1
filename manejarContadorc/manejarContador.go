// Ejecutable al que llama el servidor TCP por procesos al crear un nuevo proceso

package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//Tipo necesario para hacer los rpc
type Int int

//Variables para conexion
var ip string
var puertotcp string
var puertorpc string

func main() {

	puertocola := "9000"
	buf := new(bytes.Buffer)

	//Ingresa el argumento a "entrada"
	entrada := os.Args[1]

	//Aqui comienza la llamada a la cola

	// Codificar entrada de nuevo
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(entrada); err != nil {
		log.Println("manejarContador: Error codificando entrada: ", err)
	}

	log.Println("manejarContador: Codificacion exitosa")

	//Llamar a la cola
	log.Println("manejarContador: Llamado a cola en:", ip+":"+puertocola)
	conn, err := net.Dial("tcp", ip+":"+puertocola)
	if err != nil {
		log.Println("manejarContador: Error en dial: ", err)
	}

	defer conn.Close()

	// Escribir a la cola
	buf.WriteTo(conn)
	log.Printf("manejarContador: Buf escrito a cola")

	//Lee la conexion y la imprime
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("manejarContador: Error leyendo la conexión de regreso", err)
	} else {
		log.Printf("manejarContado: Leida la respuesta de conexion a cola con exito.")
	}
	fmt.Println(string(resp)) //Salida de manejarContador

	io.WriteString(conn, fmt.Sprintf(resp))

}
