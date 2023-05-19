package main

import (
	"fmt"
	"strings"
)

// Escreva uma função que conte a ocorrência de cada palavra em uma string e retorne um mapa onde as chaves são as
// palavras encontradas e os valores são o número de ocorrências de cada palavra.

func contarPalavras(texto string) map[string]int {
	palavras := strings.Fields(texto)
	ocorrencias := make(map[string]int)
	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}
	return ocorrencias
}

func main() {
	frase := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam eget ligula eu lectus lobortis condimentum."
	fmt.Println(contarPalavras(frase))
}
