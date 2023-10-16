package main

type Livro struct {
	Titulo string
	Autor  string
}

func mudarAutor(ptr *Livro) {
	if ptr.Autor == "Anonimo" {
		ptr.Titulo = "Desconhecido"

	}
}
