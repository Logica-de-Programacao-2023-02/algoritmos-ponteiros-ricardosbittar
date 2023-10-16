package main

import "fmt"

func modificarValor(n int, ptr *int) {

	if ptr != nil {

		*ptr = n
	} else {
		fmt.Println("ERRO: O ponteiro nao aponta para uma memoria valida")
	}
}
