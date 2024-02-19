// Golang program to illustrate the
// strings.Fields() Function
package main

import (
	"fmt"
	"strings"
)

func main() {

	// String s se divide sobre la base de espacios en blanco
	// y lo guarda en una array de cadenas
	s := "GeeksforGeeks is a computer science portal !"
	v := strings.Fields(s)
	fmt.Println(v)

	// Otro ejemplo pasando la cadena como argumento
	// directamente a la función Fields()
	v = strings.Fields("I am a software developer, I love coding")
	fmt.Println(v)
}
