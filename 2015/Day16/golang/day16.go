package main

import (
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	sues := LoadSues(ReadFile("prod.txt"))
	results := GetResults()
	fmt.Println(sues)
	fmt.Println(results)

	matches := make([]Sue,0)
	for _, sue := range sues {
		match := true
		for key, value := range sue.properties {
			result, _ := results[key]
			indicator := value == result
			// Remove switch for part 1
			switch key {
			case "cats":
				indicator = value > result
			case "trees":
				indicator = value > result
			case "pomeranians":
				indicator = value < result
			case "goldfish":
				indicator = value < result
			}
			if !indicator {
				match = false
				break
			} 
		}
		if match {
			matches = append(matches,sue)
		}
	}
	fmt.Println(matches)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}


func GetResults() map[string]int {
	r := make(map[string]int,0)
	r["children"] = 3
	r["cats"] = 7
	r["samoyeds"] = 2
	r["pomeranians"] = 3
	r["akitas"] = 0
	r["vizslas"] = 0
	r["goldfish"] = 5
	r["trees"] = 3
	r["cars"] = 2
	r["perfumes"] = 1
	return r
}

type Sue struct {
	number int
	properties map[string]int
}

func LoadSue(input string) Sue {
	regStr := `^Sue ([0-9]+): (.*)$`
	r := regexp.MustCompile(regStr)
	match := r.FindStringSubmatch(input)
	var sue Sue
	sue.properties = make(map[string]int,0)
	if len(match) == 3 {
		sueNum, _ := strconv.ParseInt(match[1],10,32)
		sue.number = int(sueNum)
		props := strings.Split(match[2],", ")
		fmt.Println(props)
		for _, kv := range props {
			prop := strings.Split(kv,": ")
			fmt.Println(prop)
			propNum, _ := strconv.ParseInt(prop[1],10,32)
			sue.properties[prop[0]] = int(propNum)
		}
	}
	return sue
}

func LoadSues(inputs []string) []Sue {
	sues := make([]Sue,0)
	for _, input := range inputs {
		sues = append(sues,LoadSue(input))
	}
	return sues
}