// Dup1 exibe o texto de toda linha que aparece mais de uma
// vez na entrada-padrão, precedida por sua contagem.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // função make() cria um map vazio (mas faz outras coisas tbm)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "" {
			break
		} else {
			counts[text]++
		}
	}
	// NOTA: ignorando erros em potencial de input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
