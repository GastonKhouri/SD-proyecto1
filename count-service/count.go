package count

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	counter int = 0 // Variable de memoria compartida
)

// Con esta funcion se escoge que operacion se quiere hacer
func realizarFuncion(function int) {
	var wg sync.WaitGroup
	wg.Add(1)

	switch function {
		case 1:
			go func() {
				reiniciar()
				wg.Done()		
			}()
			break
		case 2:
			go func() {
				incrementar(10)
				wg.Done()		
			}()
			break
		case 3:
			go func() {
				decrementar(10)
				wg.Done()		
			}()
			break
	}

	wg.Wait()
}

// Funcion que reinicia la cuenta
func reiniciar() {
	mutex.Lock()
	antes := counter
	counter = 0
	fmt.Printf("Contador antes: %d. Contador reiniciado a 0\n", antes)
	mutex.Unlock()
}

// Funcion que incrementa la cuenta
func incrementar(n int) {
	mutex.Lock()
	antes := counter
	counter += n
	fmt.Printf("Contador antes: %d. Contador incrementado: %d\n", antes, counter)
	mutex.Unlock()
}

// Funcion que decrementa la cuenta
func decrementar(n int) {
	mutex.Lock()
	antes := counter
	counter -= n
	fmt.Printf("Contador antes: %d. Contador decrementado: %d\n", antes, counter)
	mutex.Unlock()
}