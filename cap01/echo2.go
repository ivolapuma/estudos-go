// Echo2 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// com range, loop abaixo é semelhante ao for each de algumas linguagens
	// em cada laço do loop, range retorna um par de valores (indice e o valor do elemento)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
