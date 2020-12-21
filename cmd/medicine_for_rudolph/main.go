package main

import (
	"bufio"
	"os"
	"regexp"
)

func main(){
	f,_ := os.Open("/Users/aarayambeth/go_stuff/advent_of_code/cmd/medicine_for_rudolph/input.txt")
	s := bufio.NewScanner(f)
	reWord := regexp.MustCompile(`[a-zA-a]+`)
	conversions := make(map[string][]string)
	var baseWord string
	for s.Scan(){
		str := s.Text()
		if str == "" {
			s.Scan()
			baseWord = s.Text()
		} else {
			result := reWord.FindAllStringSubmatch(s.Text(), -1)
			if result != nil {
				if _,ok := conversions[result[0][0]]; !ok {
					conversions[result[0][0]] = make([]string,0)
				}
				conversions[result[0][0]] = append(conversions[result[0][0]],result[1][0])
			}
		}
	}
	println(baseWord)
}
