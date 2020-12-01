package main

import (  
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "strconv"
)

func getNumbers(numberList []int64, target int64)  (int64, int64) {
    for i := 0; i < len(numberList); i++ {
        for j := 0; j < len(numberList); j++ {
            //fmt.Println("[",numberList[i], numberList[j],"]")
            if i != j && numberList[i] + numberList[j] == target {
                return numberList[i], numberList[j]
            }
        }
    }
    return 0, 0
}

func getNumbers3(numberList []int64, target int64)  (int64, int64, int64) {
    for i := 0; i < len(numberList); i++ {
        for j := 0; j < len(numberList); j++ {
            for k := 0; k < len(numberList); k++ {
            //fmt.Println("[",numberList[i], numberList[j],"]")
                if i != j && j != k && i != k && numberList[i] + numberList[j] + numberList[k] == target {
                    return numberList[i], numberList[j], numberList[k]
                }
            }
        }
    }
    return 0, 0, 0
}

func main() {  
    fptr := flag.String("fpath", "prod.txt", "file path to read from")
    strVals := make([]int64,0)
    flag.Parse()

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    s := bufio.NewScanner(f)
    for s.Scan() {
        val, _ := strconv.ParseInt(s.Text(),10,64)
        strVals = append(strVals,val)
    }
    fmt.Println(strVals)
    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
    a, b := getNumbers(strVals,2020)
    fmt.Println(a,b,a*b)

    a, b, c := getNumbers3(strVals,2020)
    fmt.Println(a,b,c,a*b*c)
}