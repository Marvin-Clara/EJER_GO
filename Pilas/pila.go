package main

import "fmt"

//pila pasiva
type Pila struct {
	tope    *Nodo
	tamanio int
}

//PUSH
func (pila *Pila) push(valor interface{}) *Nodo {
	nuevoNodo := NuevoNodo(valor)
	nuevoNodo.Siguiente = pila.tope
	pila.tope = nuevoNodo
	pila.tamanio++
}

//pop
func (pila *Pila) pop() (interface{}, error) {
	if pila.EstaVacia() {
		return nil, fmt.Errorf("La pila esta vacia")
	}

}
func (pila *Pila) EstaVacia() bool {
	return pila.tope == nil

}
