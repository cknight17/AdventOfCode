package main

import (
	"fmt"
)

func main() {
	want := 36000000
	at := 10000000
	got := House(at)
	factor := at/2
	for {
		fmt.Println(at," => ",got," factor ",factor," off by ",want - got)
		if factor < 2 {
			fmt.Println("CLOSE")
			break
		} else if want < got {
			fmt.Println("LOWER")
			at = at - factor
			factor = factor / 2
			got = House(at)
		} else if want > got {
			at = at + factor
			got = House(at)
			fmt.Println("HIGHER")
		} else {
			fmt.Println("BINGO")
			break
		}

	}
}

func House(n int) int {
	accumulator := 0
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			accumulator += i
		}
	}
	return accumulator*10
}