package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	//"strconv"
	//"sort"
	"regexp"
)

func main() {
	// 284 too high
	blocks := ReadFile("prod.txt")
	language := GetReplacement(blocks[0])
	input := blocks[1]
	repls := ReplacementSets(input,language)
	fmt.Println(repls)
	fmt.Println(len(repls))
	tokens := MassageInput(blocks[1])
	fmt.Println(tokens)
	fmt.Println(CalcTokens(tokens))
	// total := 0
	// singles, count := SolveMultiples(tokens)
	// total += count
	// singles, count = SolveSingleFunc(singles)
	// total += count
	// singles, count = SolveMultiples(singles)
	// total += count
	// singles, count = SolveSingleFunc(singles)
	// total += count
	// singles, count = SolveMultiples(singles)
	// total += count
	// singles, count = SolveDoubleFunc(singles)
	// total += count
	// singles, count = SolveMultiples(singles)
	// total += count
	// singles, count = SolveDoubleFunc(singles)
	// total += count
	// singles, count = SolveSingleFunc(singles)
	// total += count
	// singles, count = SolveMultiples(singles)
	// total += count
	// fmt.Println(singles,total)
	// 510 too low
	// 518
	
	// 9 wrong
}

type Replacement struct {
	from string
	to string
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n\n")
}

func GetReplacement(input string) []Replacement {
	regStr := `^([a-zA-Z]+) \=\> ([a-zA-Z]+)$`
	r := regexp.MustCompile(regStr)
	replacements := make([]Replacement,0)
	rows := strings.Split(input,"\n")
	for _, row := range rows {
		fmt.Println(row)
		match := r.FindStringSubmatch(row)
		replacements = append(replacements,Replacement{ from: match[1], to: match[2] })
	}
	return replacements
}

func ReplacementSet(input string, key string, replacement string) []string {
	replacements := make([]string,0)
	//fmt.Println(key)
	//fmt.Println(replacement)
	counter := strings.Count(input,key)
	for i := 0; i < counter; i++ {
		currentStr := replaceNth(input,key,replacement,i+1)
		replacements = append(replacements,currentStr)
	}
	//fmt.Println(replacements)
	return replacements
}

func replaceNth(s, old, new string, n int) string {
    i := 0
    for m := 1; m <= n; m++ {
        x := strings.Index(s[i:], old)
        if x < 0 {
            break
        }
        i += x
        if m == n {
            return s[:i] + new + s[i+len(old):]
        }
        i += len(old)
    }
    return s
}

func Filter(inputs []string, target string) ([]string, bool) {
	filtered := make([]string,0)
	for _, item := range inputs {
		if item == target {
			return []string{item}, true
		} else if len(item) > len(target) {
			filtered = append(filtered,item)
		}
	}
	return filtered, false
}


func ReplacementSets(input string,repl []Replacement) []string {
	replLookup := make(map[string]bool)
	for _, repl := range repl {
		for _, item := range ReplacementSet(input,repl.from,repl.to) {
			//fmt.Println(input," ",repl.from," => ",repl.to," ",item)
			replLookup[item] = true
		}
	}
	finals := make([]string,0)
	for item, _ := range replLookup {
		finals = append(finals,item)
	}
	return finals
}

func ReplacementSetsSets(inputs []string,repl []Replacement) []string {
	finals := make([]string,0)
	for _, input := range inputs {
		finals = append(finals,ReplacementSets(input,repl)...)
	}
	return finals
}

func Flip(replacements []Replacement) []Replacement {
	nrepl := make([]Replacement,0)
	for _, item := range replacements {
		nrepl = append(nrepl,Replacement{ from: item.to, to: item.from })
	}
	return nrepl
}

func Diff(input string, target string) (string,string) {
	if input == target {
		return input, ""
	}
	for i := 0; i < len(input); i++ {
		if i < len(target) && input[i] != target[i] {
			return input[:i],input[i+1:len(input)]
		}
	}
	return "",""
}

type Token struct {
	value string
	terminal bool
}

func Tokenize(input string) []string {
	regStr := `[A-Z][a-z]+`
	r := regexp.MustCompile(regStr)
	return r.FindAllString(input,-1)
}

func CalcTokens(input string) int {
	numchars := len(input)
	pcount := 0
	ccount := 0
	for _, value := range input {
		if value == '(' || value == ')' {
			pcount++
		} else if value == ',' {
			ccount++
		}
	}
	return numchars - pcount - (2*ccount) - 1

}

func SolveMultiples(input string) (string, int) {
	regStr := `X{2,}`
	accumulator := 0
	r := regexp.MustCompile(regStr)
	matches := r.FindAllString(input,-1)
	for _, value := range matches {
		accumulator += len(value) - 1
	}
	return r.ReplaceAllString(input,"X"), accumulator
}

func SolveSingleFunc(input string) (string,int) {
	regStr := `X\(X\)`
	accumulator := 0
	r := regexp.MustCompile(regStr)
	matches := r.FindAllString(input,-1)
	for _, value := range matches {
		accumulator += len(value) - 1
	}
	return r.ReplaceAllString(input,"X"), accumulator
}

func SolveDoubleFunc(input string) (string,int) {
	regStr := `X\(X\,X\)`
	accumulator := 0
	r := regexp.MustCompile(regStr)
	matches := r.FindAllString(input,-1)
	for _, value := range matches {
		accumulator += len(value) - 1
	}
	return r.ReplaceAllString(input,"X"), accumulator
}

func MassageInput(input string) string {
	regStr := `[A-Z][a-z]*`
	input = strings.Replace(input,"Rn",`(`,-1)
	input = strings.Replace(input,"Y",",",-1)
	input = strings.Replace(input,"Ar",`)`,-1)
	r := regexp.MustCompile(regStr)
	input = r.ReplaceAllString(input,"X")
	return input
}

// func LookupReplacements(r []Replacement) map[string][]int {
// 	lm := make(map[string][]int,0)
// 	for index, item := range r {
// 		l, ok := lm[item.from]
// 		if ok {
// 			lm[item.from] = append(lm[item.from],l)
// 		} else {
// 			lm[item.from] = []int{l}
// 		}
// 	}
// 	return lm
// }