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
	blocks := ReadFile("prod.txt")
	language := GetReplacement(blocks[0])
	input := blocks[1]
	repls := ReplacementSets(input,language)
	fmt.Println(repls)
	fmt.Println(len(repls))
	// 510 too low
	// 518
	input2 := []string{"e"}
	target := blocks[1]
	i := 0
	for {
		var found bool
		i++
		input2 = ReplacementSetsSets(input2,language)
		input2, found = Filter(input2,target)
		if found {
			break
		} else {
			fmt.Println(i, " NOT FOUND ",len(input2))
		}
	}
	fmt.Println(input,target,len(input2))
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
		} else if len(item) < len(target) {
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

func LookupReplacements(r []Replacement) map[string][]int {
	lm := make(map[string]int,0)
	for index, item := range r {
		l, ok := lm[item.from]
		if ok {
			lm[item.from] = append(lm[item.from],l)
		} else {
			lm[item.from] = []int{l}
		}
	}
	return lm
}