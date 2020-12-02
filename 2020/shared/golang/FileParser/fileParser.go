package fileParser

import (  
    "bufio"
    //"flag"
    //"fmt"
    "log"
    "os"
	"strconv"
	"regexp"
	"sync"
)

type Day2line struct {
	Min int
	Max int
	Character string
	Input string
}

func linereader(filename string) []string {
	//fptr := flag.String("fpath", filename, "file path to read from")
	//flag.Parse()
	lines := make([]string,0)

    f, err := os.Open(filename)
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
		lines = append(lines,s.Text())
	}
	
    if err = s.Err(); err != nil {
        log.Fatal(err)
	}
	return lines
}

func day2parseline(line string) Day2line {
	day2LineFormat := "^([0-9]+)-([0-9]+) ([a-zA-Z]{1}): (.+)$"
	r, err := regexp.Compile(day2LineFormat)
	if err != nil {
		log.Fatal(err)
	}
	matches := r.FindStringSubmatch(line)
	if len(matches) < 5 {
		log.Fatal("Parse failed",r,day2LineFormat)
	}
	min, _ := strconv.ParseInt(matches[1],10,32)
	max, _ := strconv.ParseInt(matches[2],10,32)
	return Day2line{Min:int(min),Max:int(max),Character:matches[3],Input:matches[4]}
}

func day2parseline_concurrent(line string, ch chan Day2line, wg *sync.WaitGroup) {
	output := day2parseline(line)
	ch <- output
	wg.Done()
}

func Day2input(filename string) []Day2line {
	allinput := make([]Day2line,0)
	lines := linereader(filename)
	ch := make(chan Day2line, len(lines))
	var wg sync.WaitGroup
	
	for _, line := range lines {
		wg.Add(1)	
		go day2parseline_concurrent(line,ch,&wg)
	}
	wg.Wait()
	close(ch)
	for item := range ch {
		allinput = append(allinput,item)
	}
	
	return allinput
}