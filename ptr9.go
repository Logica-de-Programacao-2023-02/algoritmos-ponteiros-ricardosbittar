package main

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func desconto(ptr *Livro) {
	ptr.Preco *= 0.90
}