package main

import (
	"fmt"
	"math/rand"
)

func main() {
	oc := ordenCreciente{}
	pr := mostrarPrimos{}
	tituloOc := tituloOc{}
	tituloMP := tituloMP{}

	var opcion int

	for {
		// Mostrar el menú
		fmt.Println("**Menú**")
		fmt.Println("1. Tirar dado")
		fmt.Println("2. Salir")
		fmt.Println("-----------------")

		// Leer la opción del usuario
		fmt.Scanf("%d", &opcion)

		// Switch case para manejar la opción del usuario
		switch opcion {
		case 1:
			// Tirar el dado y mostrar el resultado
			numeroAleatorio := rand.Intn(6) + 1
			fmt.Println("**Resultado:**", numeroAleatorio)

			if numeroAleatorio%2 == 0 {
				fmt.Println("titulo: ", tituloOc.Title())
				fmt.Println(oc.Next())
			} else {
				fmt.Println("titulo: ", tituloMP.Title())
				fmt.Println(pr.Next())
			}

		case 2:
			// Salir del programa
			fmt.Println("¡Adiós!")
			return
		default:
			// Opción no válida
			fmt.Println("Opción no válida. Ingrese un número entre 1 y 2.")
		}
	}

}

type IntSequence interface {
	Next() int
	Title() string
}

type ordenCreciente struct{}
type mostrarPrimos struct{}
type tituloOc struct{}
type tituloMP struct{}

func (o ordenCreciente) Next() int {
	i := 0
	for i := 0; i < 30; i++ {
		fmt.Println(i)
	}
	return i
}

func (t tituloOc) Title() string {
	return "Listado de numeros en Orden Creciente"
}

func (t tituloMP) Title() string {
	return "Listado de numeros Primos"
}

func (mp mostrarPrimos) Next() int {

	i := 0

	for i := 2; i < 30; i++ {
		if mp.esPrimo(i) {
			fmt.Println(i)
		}
	}

	return i

}

func (mp mostrarPrimos) esPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
