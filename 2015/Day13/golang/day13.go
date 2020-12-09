package main

import (
	"regexp"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

func main() {
	inputs := ReadFile("prod.txt")
	hmap,hlist := HappinessMap(inputs)
	alist := Permutations(hlist)
	allSeating := CalculateHappinessAll(alist,hmap)
	sort.Sort(ByHappiness(allSeating))
	max := allSeating[len(allSeating)-1]
	fmt.Println(max.order,max.happiness)

	nhmap, nhlist := HappinessMap2(hmap,hlist)
	nalist := Permutations(nhlist)
	nallSeating := CalculateHappinessAll(nalist,nhmap)
	sort.Sort(ByHappiness(nallSeating))
	nmax := nallSeating[len(nallSeating)-1]
	fmt.Println(nmax.order,nmax.happiness)
}

type Seating struct {
	order []string
	happiness int
}
type ByHappiness []Seating
func (a ByHappiness) Len() int           { return len(a) }
func (a ByHappiness) Less(i, j int) bool { return a[i].happiness < a[j].happiness }
func (a ByHappiness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func ParseHappiness(input string) (string,string,int) {
	regString := `^(.+) would (lose|gain) ([0-9]+) happiness units by sitting next to (.+)\.$`
	r := regexp.MustCompile(regString)
	matches := r.FindAllStringSubmatch(input,-1)
	if (len(matches) > 0 && len(matches[0]) >= 5) {
		num, _ := strconv.ParseInt(matches[0][3],10,32)
		if matches[0][2] == "lose" {
			num *= -1
		}
		return matches[0][1], matches[0][4], int(num)
	} else {
		return "ERROR", "ERROR", 0
	}
}

func HappinessMap(inputs []string) (map[string]map[string]int,[]string) {
	happinessMap := make(map[string]map[string]int,0)
	happinessList := make(map[string]bool,0)
	for _, input := range inputs {
		from, to, amount := ParseHappiness(input)
		happinessFrom, ok := happinessMap[from]
		if !ok {
			happinessFrom = make(map[string]int,0)
			happinessMap[from] = happinessFrom
		} 
		happinessFrom[to] = amount
		happinessList[from] = true
		happinessList[to] = true
	}
	happinessListFinal := make([]string,0)
	for value, _ := range happinessList {
		happinessListFinal = append(happinessListFinal,value)
	}
	return happinessMap, happinessListFinal
}

func HappinessMap2(hmap map[string]map[string]int, list []string) (map[string]map[string]int,[]string) {
	selfScores := make(map[string]int,0)
	for _, person := range list {
		selfScores[person] = 0
	}
	for _, value := range hmap {
		value["self"] = 0
	}
	hmap["self"] = selfScores
	list = append(list,"self")
	return hmap, list
}

func CalculateHappiness(order []string,hmap map[string]map[string]int) (int, bool) {
	accumulator := 0
	for i := 0; i < len(order); i++ {
		left := order[i]
		right := order[(i+1)%len(order)]
		lr, ok := hmap[left][right]
		if !ok {
			return accumulator, false
		}
		accumulator += lr
		rl, ok := hmap[right][left]
		if !ok {
			return accumulator, false
		}
		accumulator += rl
	}
	return accumulator, true
}

func CalculateHappinessAll(orders [][]string,hmap map[string]map[string]int) []Seating {
	seating := make([]Seating,0)
	for _, order := range orders {
		happiness, ok := CalculateHappiness(order,hmap)
		if ok {
			var s Seating
			s.order = order
			s.happiness = happiness
			seating = append(seating,s)
		}
	}
	return seating
}

func Permutations(arr []string)[][]string{
    var helper func([]string, int)
    res := [][]string{}

    helper = func(arr []string, n int){
        if n == 1{
            tmp := make([]string, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}
