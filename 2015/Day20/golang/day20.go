package main

import (
	"fmt"
	// "github.com/mxschmitt/golang-combinations"
	// "math/bits"
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
	for val < 36000000 {
		num++
		val = Houses(num,10,0)
		//fmt.Println(num, ": ",val)
	}
	fmt.Println(num)
	fmt.Println(val)
	fmt.Println(ProperDivisors(val))

	fmt.Println(8)
	fmt.Println(PrimeFactors(8))
	fmt.Println(ProperDivisors(8))
	fmt.Println(All(PrimeFactors(8)))

	target := uint64(36000000)
	i := uint64(1)
	result := uint64(0)
	for result <= target {
		result = Houses(i,11,50)
		i++
	}
	fmt.Println(i,result)
	for i = 884521 - 10;i < 884521 + 10; i++ { 
		fmt.Println(i,Houses(i,11,50))
	}
	// 36191925 too high
	// 884521 too high

	// num = uint64(1)
	// val = uint64(1)
	// for val < 36000000/11 {
	// 	num++
	// 	val = SumOfProperDivisors(num) + num
	// 	fmt.Println(num, ": ",val)
	// }
	
}

func Houses(at uint64, per uint64, cap uint64) (presents uint64) {
	factors := All(PrimeFactors(at))
	for _, item := range factors {
		//fmt.Println(presents,item,at,per)
		if cap == 0 || item * cap > at {
			//fmt.Println("OK")
			presents += (item * per)
		} 
	}
	return presents
}

// All returns all combinations for a given string array.
// This is essentially a powerset of the given set except that the empty set is disregarded.
func All(set []uint64) (subsets []uint64) {
	length := uint(len(set))
	submap := make(map[uint64]bool)
	subsets = append(subsets,1)
	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset uint64 = 1
		
		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset *= set[object]
			}
		}
		// add subset to subsets
		submap[subset] = true
	}
	for key, _ := range submap {
		subsets = append(subsets,key)
	}
	return subsets
}


// // Combinations returns combinations of n elements for a given string array.
// // For n < 1, it equals to All and returns all combinations.
// func Combinations(set []string, n int) (subsets [][]string) {
// 	length := uint(len(set))

// 	if n > len(set) {
// 		n = len(set)
// 	}

// 	// Go through all possible combinations of objects
// 	// from 1 (only first object in subset) to 2^length (all objects in subset)
// 	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
// 		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
// 			continue
// 		}

// 		var subset []string

// 		for object := uint(0); object < length; object++ {
// 			// checks if object is contained in subset
// 			// by checking if bit 'object' is set in subsetBits
// 			if (subsetBits>>object)&1 == 1 {
// 				// add object to subset
// 				subset = append(subset, set[object])
// 			}
// 		}
// 		// add subset to subsets
// 		subsets = append(subsets, subset)
// 	}
// 	return subsets
// }
