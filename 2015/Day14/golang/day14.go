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
	reindeer := make([]Reindeer,0)
	for _, input := range inputs {
		r, ok := GetReindeer(input)
		if ok {
			reindeer = append(reindeer,r)
		} 
	}
	ts := make([]TimeStamp,0)
	for _, current := range reindeer {
		ts = append(ts,ProgressAt(current,2503))
	}
	sort.Sort(ByDistance(ts))
	fmt.Println(ts)
	max := ts[len(ts)-1]
	fmt.Println(max.distance)

	points := make(map[string]int)
	for _, r := range reindeer {
		points[r.name] = 0
	}

	for i := 1; i <= 2503; i++ {
		ts := make([]TimeStamp,0)
		for _, current := range reindeer {
			ts = append(ts,ProgressAt(current,i))
		}
		sort.Sort(ByDistance(ts))
		//fmt.Println(ts)
		max := ts[len(ts)-1]
		leadTime := max.distance
		for _, t := range ts {
			if t.distance == leadTime {
				points[t.reindeer.name]++
				fmt.Println("WINNER: ",i,t)
			} else {
				fmt.Println("LOSER: ",i,t)
			}
		}
	}
	fmt.Println(points)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

type Reindeer struct {
	name string
	kms int
	active int
	rest int
}

type TimeStamp struct {
	reindeer Reindeer
	time int
	distance int
	atRest bool
}
type ByDistance []TimeStamp
func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Less(i, j int) bool { return a[i].distance < a[j].distance }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetReindeer(input string) (Reindeer,bool) {
	var reindeer Reindeer
	found := false
	regString := `^([a-zA-Z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.$`
	r := regexp.MustCompile(regString)
	match := r.FindStringSubmatch(input)
	if len(match) == 5 {
		reindeer.name = match[1]
		kms, _ := strconv.ParseInt(match[2],10,32)
		reindeer.kms = int(kms)
		active, _ := strconv.ParseInt(match[3],10,32)
		reindeer.active = int(active)
		rest, _ := strconv.ParseInt(match[4],10,32)
		reindeer.rest = int(rest)
		found = true
	}
	return reindeer, found
}

func ProgressAt(reindeer Reindeer, at int) TimeStamp {
	t := TimeStamp {
		reindeer: reindeer,
		time: at,
		atRest: true,
	}
	base := at / (reindeer.active + reindeer.rest)
	distance := base * reindeer.kms * reindeer.active
	remaining := at % (reindeer.active + reindeer.rest)
	if remaining > reindeer.active {
		distance += (reindeer.active * reindeer.kms)
	} else {
		distance += (remaining * reindeer.kms)
		t.atRest = false
	}
	t.distance = distance
	return t
}