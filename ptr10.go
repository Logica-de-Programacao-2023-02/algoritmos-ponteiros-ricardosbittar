package main

import "fmt"

func primos(n int, ptr *[]int) {
	var primos []int
	for i := 2; i < n; i++ {
		if n < 2 {
			fmt.Println("Nao ha numeros primos")
		}
		if i >= 2 && i != n && n%i == 0 {

			primos = append(primos, i)
		}
	}
	*ptr = primos
}
