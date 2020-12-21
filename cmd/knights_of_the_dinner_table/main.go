package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var costs map[string]map[string]int
var people []string
var permArray []int
func processLine(input string){
	ops := strings.Split(input," ")
	// getting rid of the full stop
	ops[len(ops)-1] = ops[len(ops)-1][:len(ops[len(ops)-1])-1]
	// fmt.Printf("%v", ops)
	if _,ok := costs[ops[0]]; !ok {
		costs[ops[0]] = make(map[string]int)
		people = append(people,ops[0])
	}
	var negative int
	if ops[2] == "gain" {
		negative = 1
	} else {
		negative = -1
	}
	num, _ := strconv.Atoi(ops[3])
	costs[ops[0]][ops[10]] = negative*num
}

func permute() {
	for i:=len(permArray)-1;i>0;i--{
		// if we find a number less than current number
		if permArray[i] > permArray[i-1]{
			// swap permArray[i-1] with first number from permArray[i] to end
			// which is greater than permArray[i-1]
			for j:=len(permArray)-1;j>=i;j--{
				if permArray[i-1] < permArray[j]{
					//swap
					temp := permArray[i-1]
					permArray[i-1] = permArray[j]
					permArray[j] = temp
					break
				}
			}
			// Now make sure that permArray[i] to end of permArray is sorted in increasing order
			// note that the numbers in this range WILL be decreasing order.. so we can simply swap
			// values i.e i with len(permArray)-1, i+1 with len(permArray)-2 and so on
			first := i
			last := len(permArray)-1
			for first < last {
				if permArray[first] > permArray[last]{
					// swap
					temp:=permArray[first]
					permArray[first] = permArray[last]
					permArray[last] = temp
				}
				first++
				last--
			}
			break
		}
	}
}

func main () {
	f, _ := os.Open("/Users/aarayambeth/go_stuff/advent_of_code/cmd/knights_of_the_dinner_table/input.txt")
	s := bufio.NewScanner(f)
	costs = make(map[string]map[string]int)
	for s.Scan(){
		processLine(s.Text())
	}
	for i:=0;i<len(people);i++{
		permArray = append(permArray,i)
	}
	var limit int
	limit = 1
	for _,i := range permArray {
		limit *= i+1
	}
	println(limit)
	max := 0
	for limit > 0{
		localHappiness :=0
		for i:=0;i<len(permArray);i++{
			if i== len(permArray)-1 {
				localHappiness += costs[people[permArray[i]]][people[permArray[0]]] + costs[people[permArray[0]]][people[permArray[i]]]
			} else {
				localHappiness += costs[people[permArray[i]]][people[permArray[i+1]]] + costs[people[permArray[i+1]]][people[permArray[i]]]
			}
		}
		if localHappiness > max {
			max = localHappiness
		}
		fmt.Printf("%v",permArray)
		permute()
		limit--
	}
	println(max)
}
