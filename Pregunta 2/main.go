package main

import (
	"fmt"
)

/*
Versión modificada de la Criba de Eratóstenes, donde
en lugar de almacenar dentro de la Criba si un número
es primo o no, se almacena su factor primo más pequeño.
*/
func eratosthenesSieve(length int) []int {
	sieve := make([]int, length)
	for i := 0; i < len(sieve); i++ {
		if sieve[i] == 0 {
			sieve[i] = i + 2
			j := i
			for j+(i+2) < len(sieve) {
				j += i + 2
				if sieve[j] == 0 {
					sieve[j] = i + 2
				}
			}
		}
	}
	return sieve
}

/*
Función que calcula los factores primos de un número dada
una criba de Eratóstenes pre-calculada. Los factores primos
son retornados como un mapa donde las llaves son las bases
de los factoes y su respectivo valor es el exponente que posee
la base dentro de la repersentación en factores primos.
*/
func primeFactors(n int, es []int) map[int]int {
	pf := make(map[int]int, 0)

	for n > 1 {
		sd := es[n-2]
		_, ok := pf[sd]

		if !ok {
			pf[sd] = 1
		} else {
			pf[sd] += 1
		}
		n /= sd
	}

	return pf
}

/*
Implementación de algoritmo para el cómputo de decomp(X)
*/
func decomp(n int) int64 {
	accum := int64(0)
	es := eratosthenesSieve(n)

	for i := 1; i < n; i++ {
		pf := primeFactors(i, es)
		nd := 1
		for _, v := range pf {
			nd *= (v + 1)
		}
		accum += int64(2 * nd)
	}

	return accum
}

func main() {
	n := 24
	fmt.Printf("decomp(%d): %d\n", n, decomp(n))
}
