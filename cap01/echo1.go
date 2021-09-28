// Echo1 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// os.Args dá acesso aos argumentos de linha de comando
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
