package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 100
	fmt.Println(numero)

	var numero2 uint = 10000
	fmt.Println(numero2)

	var numero3 rune = 123456 // int32
	fmt.Println(numero3)

	var numero4 byte = 123 // int8
	fmt.Println(numero4)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 999.99
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
