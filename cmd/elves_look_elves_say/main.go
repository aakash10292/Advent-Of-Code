package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func doIt(input string) int {
	oldops := strings.Split(input,"")
	for i:=0;i<50;i++{
		var newops []string
		for j:=0;j<len(oldops);{
			cur := oldops[j]
			k := j+1
			count := 1
			for k<len(oldops) && oldops[k] == cur{
				count++
				k++
			}
			newops = append(newops,strconv.Itoa(count))
			newops = append(newops,cur)
			j = k
		}
		//fmt.Printf("%v",newops)
		//result += len(newops)
		oldops = strings.Split(strings.Join(newops,""),"")
	}
	return len(oldops)
}

func main(){
	f,_ := os.Open("/home/aarayambeth/go-stuff/advent_of_code/elves_look_elves_say/input.txt")
	s := bufio.NewScanner(f)
	var result int
	for s.Scan(){
		 result = doIt(s.Text())
	}

	println(result)
}
