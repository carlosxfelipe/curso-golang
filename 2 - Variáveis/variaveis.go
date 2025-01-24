package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "V3"
		variavel4 string = "V4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "V5", "V6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"

	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
