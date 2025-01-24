package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 100
	fmt.Println(numero)

	var numero2 uint = 18446744073709551615
	fmt.Println(numero2)

	var numero3 rune = 2147483647 // int32
	fmt.Println(numero3)

	var numero4 byte = 255 // uint8
	fmt.Println(numero4)

	var numeroReal float32 = 3.4028235e+38
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 1.7976931348623157e+308
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B' // aspas simples = tabela ASCII
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("erro aqui")
	fmt.Println(erro)
}
