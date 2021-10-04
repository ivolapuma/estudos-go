// Testando Printf()...

package main

import "fmt"

func main() {

	var s string = "10"
	fmt.Printf("s (valor)   = %v\n", s)
	fmt.Printf("s (string)  = %s\n", s)
	fmt.Printf("s (inteiro) = %d\n", s) // compila, mas dá um alerta e nao exibe o valor ao executar

	var i int = 10
	fmt.Printf("i (valor)   = %v\n", i)
	fmt.Printf("i (string)  = %s\n", i) // compila, mas dá um alerta e nao exibe o valor ao executar
	fmt.Printf("i (inteiro) = %d\n", i)
	fmt.Printf("i (hexa)    = %x\n", i)
	fmt.Printf("i (octa)    = %o\n", i)
	fmt.Printf("i (binario) = %b\n", i)
	fmt.Printf("i (float)   = %f\n", i) // compila, mas dá um alerta e nao exibe o valor ao executar

	var f32 float32 = 1.23456789
	fmt.Printf("f32 (valor)   = %v\n", f32)
	fmt.Printf("f32 (string)  = %s\n", f32) // compila, mas dá um alerta e nao exibe o valor ao executar
	fmt.Printf("f32 (inteiro) = %d\n", f32) // compila, mas dá um alerta e nao exibe o valor ao executar
	fmt.Printf("f32 (hexa)    = %x\n", f32)
	fmt.Printf("f32 (octa)    = %o\n", f32) // compila, mas dá um alerta e nao exibe o valor ao executar
	fmt.Printf("f32 (binario) = %b\n", f32)
	fmt.Printf("f32 (float)   = %f\n", f32)
	fmt.Printf("f32 (float2)  = %g\n", f32)
	fmt.Printf("f32 (float3)  = %e\n", f32)

	var f64 float64 = 1.23456789
	fmt.Printf("f64 (valor)   = %v\n", f64)
	fmt.Printf("f64 (string)  = %s\n", f64) // compila, mas dá um alerta e nao exibe o valor ao executar
	fmt.Printf("f64 (inteiro) = %d\n", f64) // compila, mas dá um alerta e nao exibe o valor ao executar
	fmt.Printf("f64 (hexa)    = %x\n", f64)
	fmt.Printf("f64 (octa)    = %o\n", f64) // compila, mas dá um alerta e nao exibe o valor ao executar
	fmt.Printf("f64 (binario) = %b\n", f64)
	fmt.Printf("f64 (float)   = %f\n", f64)
	fmt.Printf("f64 (float2)  = %g\n", f64)
	fmt.Printf("f64 (float3)  = %e\n", f64)

}
