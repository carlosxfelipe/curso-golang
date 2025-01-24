package auxiliar

import "fmt"

// Escrever exibe uma mensagem na saída padrão informando que o texto foi escrito pelo pacote auxiliar.
// Além disso, chama a função privada escrever2() para complementar sua funcionalidade.
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
