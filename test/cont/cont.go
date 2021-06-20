package cont

import "sync"

type Cuenta struct {
	mutex sync.RWMutex
	valor int
}

// funcion que aumenta o decrementa el contador dado un numero
func (c *Cuenta) Aumentar(num int) {
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
func (c *Cuenta) Obtener() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.valor
}
