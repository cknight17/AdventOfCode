package main

import (
	"aoc.knighttechnology.net/passport"
	"fmt"
)

func main() { 
	passports,results := passport.GetPassports(passport.ReadFile("prod.txt"))
	//fmt.Println(results)
	fmt.Println(passport.NumberOfValidPassports(results))
	fmt.Println(passport.NumberOfValidVerifiedPassports(passports,results))
}