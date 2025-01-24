package main

import "fmt"

func main() {
	var variavel1 string = "v1"
	variavel2 := "v2"

	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "v3"
		variavel4 string = "v4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "v5", "v6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "c1"

	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
