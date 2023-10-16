package main

func somarDigitos(ptr *int) {
	i := *ptr
	ultimo := i % 10
	i /= 10
	penultimo := i % 10
	soma := ultimo + penultimo
	*ptr = soma
}
