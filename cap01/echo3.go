// Echo3 exibe seus argumentos de linha de comando usando strings.Join()
package main

import (
	"fmt"
	"os"
	"strings"
)

func concatena_strings_com_range() {

	fmt.Println("Concatena strings com for range:")

	//s, sep := "", ""
	//outras formas de fazer a declaração de variaveis acima:
	var s string // forma recomendada
	var sep = "" // forma usada excepcionalmente (em declarações de múltiplas variáveis)

	// com range, loop abaixo é semelhante ao for each de algumas linguagens
	// em cada laço do loop, range retorna um par de valores (indice e o valor do elemento)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

}

func concatena_strings_com_Join() {
	fmt.Println("Concatena strings com strings.Join:")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	concatena_strings_com_range()
	concatena_strings_com_Join()
}
