Inteiros com sinal:

- int: tamanho dependente da arquitetura (32 ou 64 bits)
- int8: -128 a 127
- int16: -32,768 a 32,767
- int32: -2,147,483,648 a 2,147,483,647
- int64: -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807

Inteiros sem sinal:

- uint: tamanho dependente da arquitetura (32 ou 64 bits)
- uint8: 0 a 255 (alias: byte)
- uint16: 0 a 65,535
- uint32: 0 a 4,294,967,295
- uint64: 0 a 18,446,744,073,709,551,615
- uintptr: suficiente para conter o valor de um ponteiro

Números de ponto flutuante:

- float32: precisão de 6-9 dígitos decimais
- float64: precisão de 15-17 dígitos decimais

Números complexos:

- complex64: partes reais e imaginárias como float32
- complex128: partes reais e imaginárias como float64

Alias:

- byte: alias para uint8
- rune: alias para int32 (representa pontos de código Unicode)
