package main

import (
	"testing"
	"fmt"
	"sort"
)

func TestCombo2(t *testing.T) {
	ingredients := GetIngredients(ReadFile("test.txt"))
	recipes := AllRecipes(ingredients,100)
	results := CalcRecipes(recipes)
	sort.Sort(ByTotal(results))
	max := results[len(results)-1]
	fmt.Println(max)
	fmt.Println(max.total)
}

