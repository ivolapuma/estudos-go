// Echo2 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func original() {
	s, sep := "", ""
	// com range, loop abaixo é semelhante ao for each de algumas linguagens
	// em cada laço do loop, range retorna um par de valores (indice e o valor do elemento)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func exercicio() {
	// com range, loop abaixo é semelhante ao for each de algumas linguagens
	// em cada laço do loop, range retorna um par de valores (indice e o valor do elemento)
	for index, arg := range os.Args[1:] {
		fmt.Printf("indice %d valor %v\n", index, arg)
	}
}

func main() {
	original()
	exercicio()
}
