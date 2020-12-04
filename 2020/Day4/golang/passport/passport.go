package passport

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"log"
)

func ReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return string(b)
}

func getRequiredFields() map[string]bool {
	requiredFields := make(map[string]bool)
	requiredFields["byr"] = false
	requiredFields["iyr"] = false
	requiredFields["eyr"] = false
	requiredFields["hgt"] = false
	requiredFields["hcl"] = false
	requiredFields["ecl"] = false
	requiredFields["pid"] = false
	return requiredFields
}

func GetPassports(input string) ([]map[string]string, []map[string]bool) {
	passports := make([]map[string]string,0)
	requiredFieldsList := make([]map[string]bool,0)
	for _, item := range strings.Split(input,"\n\n") {
		requiredFields := getRequiredFields()
		passport := make(map[string]string,0)
		for _, element := range strings.Fields(item) {
			parts := strings.Split(element,":")
			if _, ok := requiredFields[parts[0]]; ok {
				requiredFields[parts[0]] = true
			}
			passport[parts[0]] = parts[1]
		}
		requiredFieldsList = append(requiredFieldsList,requiredFields)
		passports = append(passports,passport)
		//fmt.Println(passport)
		//fmt.Println(requiredFields)
	}
	return passports, requiredFieldsList
}

func CheckRequiredFields(fields map[string]bool) bool {
	check := true
	for _, value := range fields {
		check = check && value
	}
	return check
}

func CheckValidValues(fields map[string]string) bool {
	check := true
	for key, value := range fields {
		switch key {
			case "byr":
				check = check && checkByr(value)
			case "iyr":
				check = check && checkIyr(value)
			case "eyr":
				check = check && checkEyr(value)
			case "hgt":
				check = check && checkHgt(value)
			case "hcl":
				check = check && checkHcl(value)
			case "ecl":
				check = check && checkEcl(value)
			case "pid":
				check = check && checkPid(value)
			default:
				fmt.Printf("NOT FOUND %q value %q\n",key,value)
				//check = false
		}
	}
	return check
}

func NumberOfValidPassports(passportRequiredFields []map[string]bool) int {
	counter := 0
	for _, passportRequiredField := range passportRequiredFields {
		if CheckRequiredFields(passportRequiredField) {
			counter++
		}
	}
	return counter
}

func NumberOfValidVerifiedPassports(passportFields []map[string]string,passportRequiredFields []map[string]bool) int {
	counter := 0
	for i := 0; i < len(passportFields); i++ {
		passportRequiredField := passportRequiredFields[i]
		passportField := passportFields[i]
		if CheckRequiredFields(passportRequiredField) && CheckValidValues(passportField) {
			counter++
		}
	}
	return counter
}

func checkByr(input string) bool {
	if len(input) != 4 {
		fmt.Printf("checkByr FAIL length %q\n",input)
		return false
	}
	year, _ := strconv.ParseInt(input, 10, 32)
	if (year < 1920 || year > 2002) {
		fmt.Printf("checkByr FAIL value %q\n",input)
		return false
	}
	return true
}

func checkIyr(input string) bool {
	if len(input) != 4 {
		fmt.Printf("checkIyr FAIL length %q\n",input)
		return false
	}
	year, _ := strconv.ParseInt(input, 10, 32)
	if (year < 2010 || year > 2020) {
		fmt.Printf("checkIyr FAIL value %q\n",input)
		return false
	}
	return true
}

func checkEyr(input string) bool {
	if len(input) != 4 {
		fmt.Printf("checkEyr FAIL length %q\n",input)
		return false
	}
	year, _ := strconv.ParseInt(input, 10, 32)
	if (year < 2020 || year > 2030) {
		fmt.Printf("checkEyr FAIL value %q\n",input)
		return false
	}
	return true
}

func checkHgt(input string) bool {
	hgtLineFormat := "^([0-9]+)(in|cm)$"
	r, err := regexp.Compile(hgtLineFormat)
	if err != nil {
		log.Fatal(err)
	}
	matches := r.FindStringSubmatch(input)
	if (matches == nil) {
		fmt.Printf("checkHgt FAIL match %q\n",input)
		return false
	}
	height, _ := strconv.ParseInt(matches[1],10,32)
	measure := matches[2]
	if measure == "cm" && (height < 150 || height > 193) {
		fmt.Printf("checkHgt FAIL cm range %q\n",input)
		return false
	} else if measure == "in" && (height < 59 || height > 76) {
		fmt.Printf("checkHgt FAIL in range %q\n",input)
		return false
	}
	return true
}

func checkHcl(input string) bool {
	hclLineFormat := "^#[0-9a-z]{6}$"
	r, err := regexp.Compile(hclLineFormat)
	if err != nil {
		log.Fatal(err)
	}
	if !r.MatchString(input) {
		fmt.Printf("checkHcl FAIL %q\n",input)
		return false
	}
	return true
}

func checkEcl(input string) bool {
	eclLineFormat := "^(amb|blu|brn|gry|grn|hzl|oth)$"
	r, err := regexp.Compile(eclLineFormat)
	if err != nil {
		log.Fatal(err)
	}
	if !r.MatchString(input) {
		fmt.Printf("checkEcl FAIL %q\n",input)
		return false
	}
	return true
}

func checkPid(input string) bool {
	pidLineFormat := "^[0-9]{9}$"
	r, err := regexp.Compile(pidLineFormat)
	if err != nil {
		log.Fatal(err)
	}
	if !r.MatchString(input) {
		fmt.Printf("checkPid FAIL %q\n",input)
		return false
	}
	return true
}