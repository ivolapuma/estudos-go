// Exemplo de loops (comando for)
package main

import "fmt"

func main() {

	fmt.Println("Loop for tradicional")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Loop for tipo while")
	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

	fmt.Println("Loop for tipo while(true)")
	i = 1
	for {
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}

}
