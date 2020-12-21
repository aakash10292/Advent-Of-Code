package main

import (
	"bufio"
	"os"
	"strconv"
)
var minMap map[int]int
func numberOfCombinations(remaining []int, target int, used []int) int {
	var result,sum int
	for _,num := range used {
		sum += num
	}
	if sum == target {
		minMap[len(used)]++
		return 1
	} else if sum > target {
		return 0
	}
	for pos, _ := range remaining {
		newUsed := append(used, remaining[pos])
		newRemaining := remaining[pos+1:]
		result += numberOfCombinations(newRemaining,target,newUsed)
	}
	return result
}

func main(){
	f,_ := os.Open("/Users/aarayambeth/go_stuff/advent_of_code/cmd/no_such_thing_as_too_much/input.txt")
	s := bufio.NewScanner(f)
	var nums []int
	for s.Scan(){
		num,_ := strconv.Atoi(s.Text())
		nums = append(nums,num)
	}
	//remaining := []int{20,15,10,5,5}
	used := []int{}
	target := 150
	minMap = make(map[int]int)
	println(numberOfCombinations( nums,target,used))
	println("")
}
