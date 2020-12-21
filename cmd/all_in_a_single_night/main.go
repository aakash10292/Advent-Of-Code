package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)
type void struct{}
var member void
var distances map[string]map[string]int
var cities map[string]void
var listOfCities []string
var permArray []int
func processInput(input string){
	ops := strings.Split(input," ")
	cities[ops[0]] = member
	cities[ops[2]] = member
	if num,err := strconv.Atoi(ops[4]); err == nil {
		if distances[ops[0]] == nil {
			distances[ops[0]] = make(map[string]int)
		}
		if distances[ops[2]] == nil {
			distances[ops[2]] = make(map[string]int)
		}
		distances[ops[0]][ops[2]] = num
		distances[ops[2]][ops[0]] = num
	} else {
		log.Fatal("Error while parsing distance.")
		os.Exit(-1)
	}
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

func main(){
	f,_ := os.Open("/home/aarayambeth/go-stuff/advent_of_code/all_in_a_single_night/input.txt")
	distances = make(map[string]map[string]int)
	cities = make(map[string]void)
	s := bufio.NewScanner(f)
	for s.Scan() {
		processInput(s.Text())
	}
	listOfCities = make([]string,len(cities))
	permArray = make([]int, len(cities))
	i := 0
	for k := range cities {
		listOfCities[i] = k
		permArray[i] = i
		i++
	}
	var limit int
	limit = 1
	for _,i := range permArray{
		limit *= i+1
	}
	println(limit)
	max := 0
	for limit > 0{
		localSum := 0
		for i:=0;i<len(permArray)-1;i++{
			localSum += distances[listOfCities[permArray[i]]][listOfCities[permArray[i+1]]]
		}
		if localSum > max {
			max = localSum
		}
		//fmt.Printf("%v",permArray)
		permute()
		limit--
	}
	println(max)
}
