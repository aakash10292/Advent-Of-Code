package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main(){
	//f,_ := os.Open("/home/aarayambeth/go-stuff/advent_of_code/Matchsticks/input.txt")
	input, err := ioutil.ReadFile("/home/aarayambeth/go-stuff/advent_of_code/Matchsticks/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	codeCharacters := 0
	memoryCharacters := 0

	for _, line := range lines {
		codeCharacters += len(line)
		escaped  := strconv.Quote(line)
		memoryCharacters += len(escaped)
	}

	println(memoryCharacters - codeCharacters)
}
