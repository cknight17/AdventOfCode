package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"regexp"
	"sort"
)

func main() {
	ingredients := GetIngredients(ReadFile("prod.txt"))
	recipes := AllRecipes(ingredients,100)
	results := CalcRecipes(recipes)
	sort.Sort(ByTotal(results))
	max := results[len(results)-1]
	fmt.Println(max)
	fmt.Println(max.total)

	results = TargetCalories(500,results)
	sort.Sort(ByTotal(results))
	max = results[len(results)-1]
	fmt.Println(max)
	fmt.Println(max.total)

}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}


type Ingredient struct {
	name string
	capacity int64
	durability int64
	flavor int64
	texture int64
	calories int64
}

type Recipe struct {
	ingredients []Ingredient
	amounts []int64
}

type RecipeResults struct {
	recipe Recipe
	capacity int64
	durability int64
	flavor int64
	texture int64
	calories int64
	total int64
}
type ByTotal []RecipeResults
func (a ByTotal) Len() int           { return len(a) }
func (a ByTotal) Less(i, j int) bool { return a[i].total < a[j].total }
func (a ByTotal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetIngredient(input string) Ingredient {
	var ingredient Ingredient
	regStr := `^([A-Za-z]+): capacity ([\-0-9]+), durability ([\-0-9]+), flavor ([\-0-9]+), texture ([\-0-9]+), calories ([\-0-9]+)$`
	r := regexp.MustCompile(regStr)
	match := r.FindStringSubmatch(input)
	if len(match) == 7 {
		ingredient.name = match[1]
		capacity, _ := strconv.ParseInt(match[2],10,64)
		ingredient.capacity = capacity
		durability, _ := strconv.ParseInt(match[3],10,64)
		ingredient.durability = durability
		flavor, _ := strconv.ParseInt(match[4],10,64)
		ingredient.flavor = flavor
		texture, _ := strconv.ParseInt(match[5],10,64)
		ingredient.texture = texture
		calories, _ := strconv.ParseInt(match[6],10,64)
		ingredient.calories = calories
	}
	return ingredient
}

func AllRecipes(ingredients []Ingredient, numItems int64) []Recipe {
	ar := make([]Recipe,0)
	combos := AllCombos4(numItems)
	if len(ingredients) == 2 {
		combos = AllCombos2(numItems)
	}
	for _, combo := range combos {
		r := Recipe {
			ingredients: ingredients,
			amounts: combo,
		}
		ar = append(ar,r)
	}
	return ar
}

func CalcRecipe(recipe Recipe) RecipeResults {
	rr := RecipeResults {
		recipe: recipe,
	}
	for i := 0; i < len(recipe.ingredients); i++ {
		ingredient := recipe.ingredients[i]
		amount := recipe.amounts[i]
		rr.capacity += (ingredient.capacity * amount)
		rr.durability += (ingredient.durability * amount)
		rr.flavor += (ingredient.flavor * amount)
		rr.texture += (ingredient.texture * amount)
		rr.calories += (ingredient.calories * amount)
	}
	if rr.capacity < 0 {
		rr.capacity = 0
	}
	if rr.durability < 0 {
		rr.durability = 0
	}
	if rr.flavor < 0 {
		rr.flavor = 0
	}
	if rr.texture < 0 {
		rr.texture = 0
	}
	if rr.calories < 0 {
		rr.calories = 0
	}
	rr.total = rr.capacity * rr.durability * rr.flavor * rr.texture
	return rr
}

func CalcRecipes(recipes []Recipe) []RecipeResults {
	rr := make([]RecipeResults,0)
	for _, recipe := range recipes {
		rr = append(rr,CalcRecipe(recipe))
	}
	return rr
}

func GetIngredients(inputs []string) []Ingredient {
	ingredients := make([]Ingredient,0)
	for _, input := range inputs {
		ingredients = append(ingredients,GetIngredient(input))
	}
	return ingredients
}

func AllCombos2(numItems int64) [][]int64 {
	allBuckets := make([][]int64,0)
	
	for i := 0; i <= int(numItems); i++ {
		for j := 0; j <= int(numItems); j++ {
			if i+j == int(numItems) {
				bucket := make([]int64,2)
				bucket[0] = int64(i)
				bucket[1] = int64(j)
				allBuckets = append(allBuckets,bucket)
			}

		}
	}
	return allBuckets
}

func AllCombos4(numItems int64) [][]int64 {
	allBuckets := make([][]int64,0)
	
	for i := 0; i <= int(numItems); i++ {
		for j := 0; j <= int(numItems); j++ {
			for k := 0; k <= int(numItems); k++ {
				for l := 0; l <= int(numItems); l++ {
					if i+j+k+l == int(numItems) {
						bucket := make([]int64,4)
						bucket[0] = int64(i)
						bucket[1] = int64(j)
						bucket[2] = int64(k)
						bucket[3] = int64(l)
						allBuckets = append(allBuckets,bucket)
					}
				}
			}
		}
	}
	return allBuckets
}

func TargetCalories(target int64, rr []RecipeResults) []RecipeResults {
	nrr := make([]RecipeResults,0)
	for _,r := range rr {
		if r.calories == target {
			nrr = append(nrr,r)
		}
	}
	return nrr
}