package main

import "fmt"

func PesquisaBinaria(lista []int, item int) int {
	baixo := 0
	alto := len(lista) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]

		if chute == item {
			return meio
		}

		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1
}

func main() {
	lista := []int{1, 3, 5, 7, 9}
	item := 7
	fmt.Printf("Numero encontrado na posição: %d\n", PesquisaBinaria(lista, item))
}
