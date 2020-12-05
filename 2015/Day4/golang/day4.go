package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(leading5("yzbqklnj"))
	fmt.Println(leading6("yzbqklnj"))
}

func leading5(input string) int64 {
	var at int64 = 0
	for {
		data := []byte(input + strconv.FormatInt(at,10))
		hash := md5.Sum(data)
		hashstr := hex.EncodeToString(hash[:])
		leadingZeroes := true
		for i := 0; i < 5; i++ {
			leadingZeroes = leadingZeroes && (hashstr[i] == '0')
		}
		if (leadingZeroes) {
			return at
		}
		at++
	}
	return -1
}

func leading6(input string) int64 {
	var at int64 = 0
	for {
		data := []byte(input + strconv.FormatInt(at,10))
		hash := md5.Sum(data)
		hashstr := hex.EncodeToString(hash[:])
		leadingZeroes := true
		for i := 0; i < 6; i++ {
			leadingZeroes = leadingZeroes && (hashstr[i] == '0')
		}
		if (leadingZeroes) {
			return at
		}
		at++
	}
	return -1
}