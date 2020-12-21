package main

import (
	"bufio"
	"os"
	"strings"
)

func validate(input string) bool {
	if strings.ContainsAny(input,"iol"){
		return false
	}
	chars := []byte(input)
	var found bool
	for i:=0;i<len(chars)-2;i++{
		if chars[i]+1 == chars[i+1] &&  chars[i]+2 == chars[i+2]{
			found = true
			break
		}
	}
	if !found {
		return false
	}
	pairs := 0
	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == chars[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}

func nextPassword(input string) string {
	var valid bool
	result := input
	for !valid {
		chars := []byte(result)
		i := len(chars)-1
		for i>0 && chars[i] == 'z' {
			chars[i] = 'a'
			i--
		}
		chars[i]++
		result = string(chars)
		valid = validate(result)
	}
	return result
}

func main(){
	f,_ := os.Open("/home/aarayambeth/go-stuff/advent_of_code/corporate_policy/input.txt")
	s := bufio.NewScanner(f)
	for s.Scan(){
		println(nextPassword(s.Text()))
	}
}
