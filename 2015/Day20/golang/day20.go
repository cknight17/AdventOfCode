package main

import (
	"fmt"
)

// Get all prime factors of a given number n
func PrimeFactors(n uint64) (pfs []uint64) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := uint64(3); i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// return p^i
func Power(p, i uint64) uint64 {
	result := uint64(1)
	for j := uint64(0); j < i; j++ {
		result *= p
	}
	return result
}

// formula comes from https://math.stackexchange.com/a/22723
func SumOfProperDivisors(n uint64) uint64 {
	pfs := PrimeFactors(n)

	// key: prime
	// value: prime exponents
	m := make(map[uint64]uint64)
	for _, prime := range pfs {
		_, ok := m[prime]
		if ok {
			m[prime] += 1
		} else {
			m[prime] = 1
		}
	}

	sumOfAllFactors := uint64(1)
	for prime, exponents := range m {
		sumOfAllFactors *= (Power(prime, exponents+1) - 1) / (prime - 1)
	}
	return sumOfAllFactors - n
}

func ProperDivisors(n uint64) map[uint64]uint64 {
	pfs := PrimeFactors(n)

	// key: prime
	// value: prime exponents
	m := make(map[uint64]uint64)
	for _, prime := range pfs {
		_, ok := m[prime]
		if ok {
			m[prime] += 1
		} else {
			m[prime] = 1
		}
	}

	return m
}


func main() {
	num := uint64(1)
	val := uint64(1)
	for val < 36000000/10 {
		num++
		val = SumOfProperDivisors(num) + num
		fmt.Println(num, ": ",val)
	}
	fmt.Println(ProperDivisors(val))

	// num = uint64(1)
	// val = uint64(1)
	// for val < 36000000/11 {
	// 	num++
	// 	val = SumOfProperDivisors(num) + num
	// 	fmt.Println(num, ": ",val)
	// }
	
}
