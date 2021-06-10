// En este ejemplo se tinen dos funciones con for en el que
// Una aumenta 10mil veces el contador y la otra decrementa 10mil veces

package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	counter int // Variable de memoria compartida
	loop int = 10000
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go hilo1(&wg)
	go hilo2(&wg)
	wg.Wait()
	
	fmt.Print("El valor de counter es: ")
	fmt.Println(counter)
}

func hilo1 (wg *sync.WaitGroup) {
	for i := 0; i < loop; i++ {
		mutex.Lock()
		counter += 1
	 	mutex.Unlock()
	}
	wg.Done()
}

func hilo2 (wg *sync.WaitGroup) {
	for i := 0; i < loop; i++ {
		mutex.Lock()
		counter -= 1
	 	mutex.Unlock()
	}
	wg.Done()
}
