package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	//"regexp"
	//"math"
)

type Game struct {
	values []uint64
	last map[uint64]uint64
}

func CopyGame(g Game) Game {
	nValue := make([]uint64,len(g.values))
	nLast := make(map[uint64]uint64,0)
	for key, value := range g.last {
		nLast[key] = value
	}
	//at := g.at
	return Game{values:nValue,last:nLast}
}

func main() {
	input := "15,5,1,4,7,0"
	g := LoadMap(input)
	fmt.Println("INIT: ",g)
	g = RunGame(g,2020)
	got := g.values[len(g.values)-1]
	fmt.Println(g,got)

	g = LoadMap(input)
	fmt.Println("INIT: ",g)
	g = RunGame(g,30000000)
	got = g.values[len(g.values)-1]
	fmt.Println(got)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func LoadMap(input string) Game {
	mmap := make(map[uint64]uint64)
	mlist := make([]uint64,0)
	at := 0
	pieces := strings.Split(input,",")
	for index, element := range pieces {
		value, err := strconv.ParseInt(element,10,64)
		if err != nil {
			fmt.Println("ERROR ",element)
		}
		at = index
		mlist = append(mlist,uint64(value))
		if at != len(pieces) - 1 {
			mmap[uint64(value)] = uint64(at)
		}
	}
	return Game{values:mlist,last:mmap}
}

func RunGame(g Game, rounds uint64) Game {
	at := uint64(len(g.values))
	for i := at; i < rounds; i++ {
		//fmt.Println(g)
		// Check for last value
		// If found current index value = index-1 + last index
		// If not current index 0
		// store last in map at index-1
		currentIndex := i - 1
		lastNumber := g.values[currentIndex]
		newValue := uint64(0)
		if lastIndex, ok := g.last[lastNumber]; ok {
			newValue = currentIndex - lastIndex
			//fmt.Println(i,": ",i, " - ",lastIndex," = ",newValue," where lastIndex=g.last[",lastNumber,"] = ",lastIndex," ",g.last)
		} else {
			//fmt.Println(i,": NOT FOUND ", g.values[currentIndex] ," 0 ",g.last)
		}
		g.values = append(g.values,newValue)
		g.last[lastNumber] = currentIndex
		
	}
	return g
}