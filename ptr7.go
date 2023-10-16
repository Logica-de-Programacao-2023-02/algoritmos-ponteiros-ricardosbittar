package main

type Conta struct {
	Saldo   float64
	Titular string
}

func addSaldo(ptr *Conta) {
	ptr.Saldo += 3000
}
