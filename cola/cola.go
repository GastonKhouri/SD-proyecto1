package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/process"
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

		// canalConexion <- conn

		//Decodificar el request para agregarlo a la cola como string
		dec := gob.NewDecoder(conn)
		if err = dec.Decode(&entrada); err != nil {
			log.Println("Cola: Error decodificando: ", err)
		}

		agregarACola(&solicitud{entrada, conn}) //Agregando a cola
	}
}

func enviarReqAContador(p chan (bool), canalConexion chan (net.Conn), puertorpc string) {
	for {
		time.Sleep(time.Second * 10)
		if len(cola) != 0 {
			fmt.Println("----- COLA ACTUAL -----")
			for i := 0; i < len(cola); i++ {
				fmt.Printf("Addr: %s - Entrada: %s\n", cola[i].conn.LocalAddr(), cola[i].num)
			}
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
		} else {
			manejarAumento(conn, num, ps)
		}
	} else if entrada == "r" {
		//Hubo error convirtiendo por lo tanto es 'r'
		manejarReseteo(conn, ps)
	} else if entrada == "p" {

		//Lista de procesos para filtrar
		listaDeProcesos := ([]string{"counter", "colatcp", "manejarContador", "servertcphilos", "servertcpproc", "serverudp",
			"clientudp", "clienttcpproc", "clienttcphilos", "clienterpc"})

		miscStat, _ := load.Misc()
		log.Printf("No. procesos corriendo: %d\n", miscStat.ProcsRunning)
		procesos, _ := process.Processes()

		resp := new(bytes.Buffer)

		for _, v := range procesos {
			name, _ := v.Name()
			pid, _ := v.Ppid()
			if stringInSlice(name, listaDeProcesos) {
				resp.WriteString("/PID: " + strconv.Itoa(int(pid)) + " - Nombre: " + name)
			}
		}
		resp.WriteString("\n")

		fmt.Println("La cola mando esto:", resp)
		io.WriteString(conn, fmt.Sprintf("Los procesos corriendo son: %s", resp))

	} else {
		log.Println("Ingreso una entrada invalida.")
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

func stringInSlice(s string, a []string) bool {
	for _, b := range a {
		if b == s {
			return true
		}
	}
	return false
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
