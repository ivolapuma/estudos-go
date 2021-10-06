// Exercicio 1.4: Modifique dup2 para que exiba os nomes de todos os arquivos
// em que cada linha duplicada ocorre

// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma vez na
// entrada. Ele lê de stdin ou de uma lista de arquivos nomeados.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			// se err == nil, arquivo foi aberto com sucesso
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// conta as linhas repetidas no arquivo
// parametro counts é um mapa, passado por referência a countLines
func countLines(f *os.File, counts map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if text == "" {
			break
		} else {
			text = filename + ":" + text
			counts[text]++
		}
	}
}
