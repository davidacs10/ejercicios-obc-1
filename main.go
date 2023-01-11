package main

import "fmt"

type Coche struct {
	puertas int
}

func main() {
	resultado := suma(10, 20, 30)
	fmt.Println("Suma de 3 patrones", resultado)

	var c Coche
	c.puertas = 4
	c.addDoor()
	fmt.Println("Puertas =", c.puertas)
}

func suma(a, b, c int) int {
	return a + b + c
}

func (c *Coche) addDoor() {
	(*c).puertas += 1
}
