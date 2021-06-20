//Cliente UDP ; puerto 2002

package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {

	ip := "localhost"
	puerto := "2002"
	service := ip + ":" + puerto

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Println("Cliente UDP: Error resolviendo: ", err)
	}

	for {

		var entrada string

		//Pedir entrada hasta que se ingrese un entero o 'r'
		for {
			fmt.Scanln(&entrada)

			if _, err := strconv.Atoi(entrada); err != nil {
				//No es un int, revisa si es reset
				if entrada == "r" {
					// Es un reset
					break
				} else if entrada == "p" {
					// Es una llamada a ver los procesos
					break
				} else {
					fmt.Println("Ingrese un numero entero o 'r'")
				}
			} else {
				// Es un int
				break
			}

		}

		//Llamar al servidor
		conn, err := net.DialUDP("udp4", nil, udpAddr)
		if err != nil {
			log.Println("Error en dial: ", err)
		}

		log.Println(fmt.Sprintf("Cliente UDP: Direccion remota escrita: %s", conn.RemoteAddr().String()))

		//Asegura que se cierre la conexión
		defer conn.Close()

		msj := []byte(entrada)

		conn.Write(msj)

		//Lee la conexion y la imprime
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		log.Print(fmt.Sprintln("Cliente UDP: Direccion remota leida: ", addr))
		resp := string(buffer[:n])
		resp = strings.Replace(resp, "/", "\n", -1)
		log.Print(fmt.Sprintln("Cliente UDP: Recibido: ", resp))
	}
}
