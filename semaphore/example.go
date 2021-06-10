// En este ejemplo se modifica la variable de memoria compartida
// En dos funciones distintas, aumentar y decrementar

// Se debe bloquear el mutex cuando se vaya a modificar el area critica
// Y desbloquear el mutex cuando termine su operacion
// El area critica no es mas que el momento en el que se modifica
// La variable en memoria compartida

package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	counter int // Variable de memoria compartida
)

// Funcion para aumentar contador
func aumentar(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Aumentando %d a la cuenta que actualmente es %d\n", value, counter)
	counter += value // Momento critico
	mutex.Unlock()
	wg.Done()
}

func decrementar(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Decrementando %d a la cuenta que actualmente es %d\n", value, counter)
	counter -= value // Momento critico
	mutex.Unlock()
	wg.Done()
}

func main() {
	counter = 0
	var wg sync.WaitGroup

	// No se para que es esta vaina pero hay que hacerlo
	wg.Add(2)

	// Ejecutar funciones con goroutines
	go decrementar(700, &wg)
	go aumentar(500, &wg)

	// Esto aguarda a que hayan terminado todos los procesos con el wg
	wg.Wait()

	fmt.Printf("El nuevo valor de la cuenta es %d\n", counter)
}