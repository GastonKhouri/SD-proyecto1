package main

import (
	"fmt"
	"time"
)

//Estructura de datos para quienes editen el contador
type Editor struct {
	ip     string
	numero int
	hora   time.Time
}

//Crear nuevo editor
func nuevoEditor(ip string, numero int) *Editor {
	p := Editor{ip: ip, numero: numero, hora: time.Now()}
	return &p
}

//Agregar un editor a la lista
func agregarALista(ip string, n int) {
	listaEditores = append(listaEditores, *nuevoEditor(ip, n))
}

//Imprimir los editores actuales de la lista
func imprimirEditores() {
	for _, v := range listaEditores {
		fmt.Printf("IP: %v  |  %v  |  n %v \n", v.ip, v.hora.Format("15:04:05 2006-02-01"), v.numero)
	}
}

var listaEditores = make([]Editor, 0)
