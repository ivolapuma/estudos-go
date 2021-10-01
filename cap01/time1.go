// Usando pacote "time"...
package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	for i := 1; i < 10000000000; i++ {
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Printf("t1: %v\n", t1)
	fmt.Printf("t2: %v\n", t2)
	fmt.Printf("diferença: %v\n", diff)
}
