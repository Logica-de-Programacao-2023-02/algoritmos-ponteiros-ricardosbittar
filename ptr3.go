package main

func reverso(ptr *string) {
	letras := []rune(*ptr)
	for i, j := 0, len(letras)-1; i < j; i, j = i+1, j-1 {
		letras[i], letras[j] = letras[j], letras[i]
		*ptr = string(letras)
	}
}
