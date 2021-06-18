// package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/http"
// 	"net/rpc"
// 	"strconv"
// )

// func (a *API) AgregarACola(n String, resp *String) error {
// 	log.Printf("Cola: Ejecutando procedimiento agregarACola con: %s\n", n)
// 	cola = append(cola, string(n))
// 	*resp = n

// 	return nil
// }

// func dequeue() string {
// 	element := cola[0]
// 	cola = cola[:1]
// 	return element
// }

// //Tipo necesario para hacer los rpc
// type String string
// type Int int
// type API int

// var cola []string
// var puertoentrada string
// var puertosalida string
// var ip string

// func main() {

// 	var num int32
// 	puertoentrada = "9000"
// 	puertosalida = "9001"
// 	ip = "localhost"
// 	var resp string
// 	var api = new(API)
// 	err := rpc.Register(api)
// 	if err != nil {
// 		log.Println("Cola: Error en registrar api: ", err)
// 	}

// 	//Recibir numero de parte de los servidores
// 	rpc.HandleHTTP()

// 	//Escuchar
// 	ls, err := net.Listen("tcp", ip+":"+puertoentrada)
// 	if err != nil {
// 		log.Println("Cola: Error escuchando (listen): ", err)
// 	}

// 	//Sirviendo
// 	log.Printf("Cola: sirviendo RPC en puerto %s", puertoentrada)
// 	err = http.Serve(ls, nil)
// 	if err != nil {
// 		log.Fatal("Counter: error sirviendo: ", err)
// 	}

// 	//Envio al contador para ejecutar el siguiente numero
// 	for {
// 		//Se saca el primer elemento
// 		entrada := dequeue()

// 		//Intenta convertir la entrada a int, si no puede, la entrada es r entonces resetea
// 		if x, err := strconv.Atoi(entrada); err == nil {
// 			//No hubo error convirtiendo porque es int
// 			num = int32(x)
// 			if num == 0 {
// 				manejarValor(puertosalida, &resp)
// 				log.Println("Se manejo valor en cola")
// 			} else {
// 				manejarAumento(num, puertosalida, &resp)
// 				log.Println("Se manejo aumento en cola")
// 			}
// 		} else {
// 			//Hubo error convirtiendo por lo tanto es 'r'
// 			manejarReseteo(puertosalida, &resp)
// 			log.Println("Se manejo reseteo en cola")
// 		}
// 		if &resp != nil {
// 			log.Println("XXX LOOP")
// 			continue
// 		}
// 	}

// }

// //Funciones para llamar los RPC
// func manejarAumento(n int32, puerto string, respuesta *string) {

// 	var resp Int

// 	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

// 	if err != nil {
// 		log.Fatal("Error de conexión: ", err)
// 	}

// 	client.Call("API.Aumentar", n, &resp)
// 	log.Println(fmt.Sprintf("Cola: Hice una llamada para aumentar y devolvió %d\n", resp)) //Log
// 	*respuesta = fmt.Sprintf("Hice una llamada para aumentar y devolvió %d \n", resp)      //Respuesta que ira al servidor

// 	client.Close()

// }

// func manejarValor(puerto string, respuesta *string) {
// 	var resp Int

// 	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

// 	if err != nil {
// 		log.Fatal("Server UDP: Error de conexión: ", err)
// 	}

// 	client.Call("API.Valor", 0, &resp)
// 	log.Println(fmt.Sprintf("Server UDP: El valor actual es: %d\n", resp))
// 	*respuesta = fmt.Sprintf("El valor actual es: %d \n", resp)
// 	client.Close()
// }

// func manejarReseteo(puerto string, respuesta *string) {
// 	var resp Int

// 	client, err := rpc.DialHTTP("tcp", ip+":"+puerto)

// 	if err != nil {
// 		log.Fatal("Server UDP: Error de conexión", err)
// 	}

// 	client.Call("API.Reset", 0, &resp)
// 	log.Println(fmt.Sprintf("Server UDP: Contador reseteado. Valor actual: %d\n", resp))
// 	*respuesta = fmt.Sprintf("Contador reseteado. Valor actual: %d \n", resp)
// 	client.Close()
// }
