package main

func update(n int, ptr *int) int {
	var ptr2 *int
	var soma int
	for i := 0; i < n; i++ {
		soma += i

	}
	*ptr = n
	*ptr2 = soma
	*ptr, ptr2 = *ptr2, *ptr
}
func main() {
	var ptr *int
	n := 10

}