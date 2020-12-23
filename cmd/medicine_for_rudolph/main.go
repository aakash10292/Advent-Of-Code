package main

import (
	"bufio"
	"os"
	"regexp"
)

func replaceInString(startIndex int, endIndex int, baseWord string, replacement string) string {
	out := []rune(baseWord)
	in := []rune(replacement)
	result := []rune{}
	for i:=0;i<startIndex;i++{
		result = append(result, out[i])
	}
	result = append(result,in...)
	result = append(result,out[endIndex:]...)

	//val := string(result)
	//println(val)
	return string(result)
}

func main(){
	f,_ := os.Open("/Users/aarayambeth/go_stuff/advent_of_code/cmd/medicine_for_rudolph/input.txt")
	s := bufio.NewScanner(f)
	type void struct{}
	var member void
	reWord := regexp.MustCompile(`[a-zA-a]+`)
	conversions := make(map[string][]string)
	result := make(map[string]void)
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
	for k,v := range conversions {
		reKey := regexp.MustCompile(k)
		matches := reKey.FindAllStringSubmatchIndex(baseWord,-1)
		for i:=0;i<len(matches);i++{
			// For each of the matches found
			for j:=0;j<len(v);j++{
				// for each possible conversion associated with this molecule,
				// replace and create new word
				newBaseWord := replaceInString(matches[i][0],matches[i][1],baseWord,v[j])
				_,ok:=result[newBaseWord]; if !ok {
					result[newBaseWord] = member
				}
			}
		}
	}

	println(len(result))
}
