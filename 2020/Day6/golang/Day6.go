package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	groups := GetGroups(ReadFile("test.txt"))
	for _, group := range groups {
		fmt.Println(CountAnswers(group))
	}

	ngroups := GetGroups(ReadFile("prod.txt"))
	tally := 0
	for _, group := range ngroups {
		tally += CountAnswers(group)
	}
	fmt.Println(tally)
	ngroups = GetGroups2(ReadFile("test.txt"))
	tally = 0
	for _, group := range ngroups {
		tally += CountAllAnswers(group)
	}
	fmt.Println(tally)
	ngroups = GetGroups2(ReadFile("prod.txt"))
	tally = 0
	for _, group := range ngroups {
		tally += CountAllAnswers(group)
	}
	fmt.Println(tally)
}

func ReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return string(b)
}

func GetGroups(input string) []string {
	groups := strings.Split(input,"\n\n")
	ngroups := make([]string,0)
	for _,group := range groups {
		ngroups = append(ngroups,strings.Replace(group,"\n","",-1))
	}
	return ngroups
}

func CountAnswers(input string) int {
	inputs := make(map[rune]bool,0)
	for _, item := range input {
		inputs[item] = true
	}
	return len(inputs)
}

func GetGroups2(input string) []string {
	groups := strings.Split(input,"\n\n")
	return groups
}

func CountAllAnswers(input string) int {
	people := strings.Split(input,"\n")
	totalPeople := len(people)
	answers := make(map[rune]int,0)
	for _, person := range people {
		for _, answer := range person {
			value, ok := answers[answer]
			if !ok {
				value = 0
			}
			answers[answer] = value + 1
		}
	}
	counter := 0
	for _, item := range answers {
		if item == totalPeople {
			counter++
		}
	}
	return counter
}