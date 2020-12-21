package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	//"sort"
	"regexp"
)

func main() {
	tt := ParseSections(ReadFile("prod.txt"))
	tt = FindErrors(tt)
	errorRate := TicketScanningErrorRate(tt)
	fmt.Println("Error rate: ",errorRate)

	tt = ParseSections(ReadFile("prod.txt"))
	tt = FindErrors(tt)
	tt = ApplyDesignations(tt)
	tt = ValidateDesignations(tt)
	fmt.Println(tt.myFinalTicket)
	departure := int64(1)
	for key, value := range tt.myFinalTicket {
		if strings.HasPrefix(key,"departure") {
			departure *= value
		}
	}
	fmt.Println("Deprarture product: ",departure)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

type TicketTranslation struct {
	designations map[int64][]string
	myTicket []int64
	tickets [][]int64
	validTickets [][]int64
	errors []int64
	errorTicket map[int64]bool
	ticketsWithDesignations [][][]string
	validDesignations []string
	myFinalTicket map[string]int64
}

func ParseSections(inputs []string) (tt TicketTranslation) {
	tt.tickets = make([][]int64,0)
	tt.designations = make(map[int64][]string,0)
	regStrDesignation := `([a-z\s]+)\: ([\-0-9]+) or ([\-0-9]+)`
	r := regexp.MustCompile(regStrDesignation)
	yourTicket := "your ticket:"
	nearbyTickets := "nearby tickets:"

	yourTicketMode := false
	nearbyTicketMode := false
	for _, input := range inputs {
		if matches := r.FindStringSubmatch(input); matches != nil {
			name := matches[1]
			range1 := matches[2]
			range2 := matches[3]
			fullRange := parseRange(range1)
			fullRange = append(fullRange,parseRange(range2)...)
			for _, index := range fullRange {
				entry, ok := tt.designations[index]
				if ok {
					tt.designations[index] = append(entry,name)
				} else {
					tt.designations[index] = []string{name}
				}
			}
		} else if input == yourTicket {
			yourTicketMode = true
			nearbyTicketMode = false
		} else if input == nearbyTickets {
			yourTicketMode = false
			nearbyTicketMode = true
		} else if yourTicketMode {
			if len(input) > 0 {
				numberStrs := strings.Split(input,",")
				tt.myTicket = make([]int64,0)
				for _, numStr := range numberStrs {
					num, err := strconv.ParseInt(numStr,10,64)
					if err == nil {
						tt.myTicket = append(tt.myTicket,num)
					} else {
						fmt.Println("CAN'T PARSE: ",numStr,err)
					}
				}
			}
		} else if nearbyTicketMode {
			if len(input) > 0 {
				numberStrs := strings.Split(input,",")
				ticket := make([]int64,0)
				for _, numStr := range numberStrs {
					num, err := strconv.ParseInt(numStr,10,64)
					if err == nil {
						ticket = append(ticket,num)
					} else {
						fmt.Println("CAN'T PARSE: ",numStr,err)
					}
				}
				tt.tickets = append(tt.tickets,ticket)
			}
		}
	}
	return tt
}

func parseRange(rangen string) []int64 {
	fullRange := make([]int64,0)
	elements := strings.Split(rangen,"-")
	min, _ := strconv.ParseInt(elements[0],10,64)
	max, _ := strconv.ParseInt(elements[1],10,64)
	for i := min; i <= max; i++ {
		fullRange = append(fullRange,i)
	}
	return fullRange
}

func FindErrors(tt TicketTranslation) TicketTranslation {
	tt.errors = make([]int64,0)
	tt.errorTicket = make(map[int64]bool,0)
	for ticketIndex, ticket := range tt.tickets {
		for _, designation := range ticket {
			_, ok := tt.designations[designation]
			if !ok {
				tt.errors = append(tt.errors,designation)
				tt.errorTicket[int64(ticketIndex)] = true
			}
		}
	}

	tt.validTickets = make([][]int64,0)
	for index, ticket := range tt.tickets {
		_, ok := tt.errorTicket[int64(index)]
		if !ok {
			tt.validTickets = append(tt.validTickets,ticket)
		}
	}
	return tt
}

func TicketScanningErrorRate(tt TicketTranslation) int64 {
	sum := int64(0)
	for _, input := range tt.errors {
		sum += input
	}
	return sum
}

func ApplyDesignations(tt TicketTranslation) TicketTranslation {
	tt.ticketsWithDesignations = make([][][]string,0)
	

	for _, validTicket := range tt.validTickets {
		ticketFieldLength := len(validTicket)
		ticketsWithDesignation := make([][]string,0)
		for i := 0; i < ticketFieldLength; i++ {
			ticketsWithDesignation = append(ticketsWithDesignation,tt.designations[validTicket[i]])
		}
		tt.ticketsWithDesignations = append(tt.ticketsWithDesignations,ticketsWithDesignation)
	}
	//fmt.Println(tt)
	return tt
}

func ValidateDesignations(tt TicketTranslation) TicketTranslation {
	length := int64(len(tt.ticketsWithDesignations[0]))
	//fmt.Println(length)
	target := int64(len(tt.ticketsWithDesignations))
	tt.validDesignations = make([]string,length)
	mappings := make([]map[string]bool,0)
	for i := int64(0); i < length; i++ {
		mapping := make(map[string]int64,0)
		for _, validTicket := range tt.ticketsWithDesignations {
			currentDesignations := validTicket[i]
			for _, designation := range currentDesignations {
				if _, ok := mapping[designation]; ok {
					mapping[designation]++
				} else {
					mapping[designation] = 1
				}
			}
		}
		bmapping := make(map[string]bool,0) 
		for designation, num := range mapping {
			bmapping[designation] = num == target
		}
		mappings = append(mappings,bmapping)
	}
	tt.validDesignations = DeriveDesignations(mappings)

	tt.myFinalTicket = make(map[string]int64,len(tt.myTicket))
	//fmt.Println(tt.myTicket,len(tt.myTicket))
	//fmt.Println(tt.validDesignations)
	for i := 0; i < len(tt.myTicket); i++ {
		//fmt.Printf("tt.myFinalTicket[%q] = %d\n",tt.validDesignations[i],tt.myTicket[i])
		tt.myFinalTicket[tt.validDesignations[i]] = tt.myTicket[i]
	}

	//fmt.Println(tt.myFinalTicket)
	return tt
}

func DeriveDesignations(mappings []map[string]bool) (designations []string) {
	//fmt.Println(mappings)
	target := len(mappings)
	designations = make([]string,target)
	found := 0
	foundMap := make(map[int64]bool)
	foundDesignation := make(map[string]bool)
	
	for found < target {
		lastFound := found
		//fmt.Printf("Found: %d Target: %d\n",found,target)
		//fmt.Println(designations)
		for index, _ := range mappings {
			if found, ok := foundMap[int64(index)]; found && ok {
				//fmt.Printf("Found index %d\n",index)
				continue
			} 
			
			validCount := 0
			var validItem string
			possibles := make([]string,0)
			for designation, valid := range mappings[index] {
				if found, ok := foundDesignation[designation]; !(found && ok) {
					if valid {
						validCount++
						validItem = designation
						possibles = append(possibles,designation)
					}
				}
			}
			if validCount == 1 {
				designations[index] = validItem
				foundMap[int64(index)] = true
				foundDesignation[validItem] = true
				found++
				//fmt.Println(index," FOUND ",validItem)
			} else if validCount > 0 {
				//fmt.Println(index," could be ",possibles)
			} else {
				//fmt.Println("ERROR, none remain")
			}
		}
		if found == lastFound {
			//fmt.Println("Uh oh")
		}
	}
	//fmt.Printf("Found: %d Target: %d\n",found,target)
	return designations
}


