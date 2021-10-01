// Testando pacote "bufio"

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("os.Stdin: %v\n", os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("scanner: %v\n", scanner)
	//fmt.Printf("scanner.Scan(): %v\n", scanner.Scan())
	x := scanner.Scan()
	y := scanner.Text()
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("(y==\\r\\n): %v\n", (y == "\r\n"))
	fmt.Printf("Compare(y, \\r\\n): %v\n", strings.Compare(y, "\r\n"))
	fmt.Printf("(y==''): %v\n", (y == ""))
	fmt.Printf("Compare(y, ''): %v\n", strings.Compare(y, ""))
	// ao digitar Enter após Scan(), y recebe "" de Text()
}
