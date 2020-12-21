package main

import (
	"encoding/json"
	"io/ioutil"
)

func processInput(input interface{}) []int {
	var numbers []int
	switch inp := input.(type){
	case []interface{}:
		for _,value := range inp{
			numbers = append(numbers, processInput(value)...)
		}
	case map[string]interface{}:
		noRed := true

		for _, value := range inp {
			if str, ok := value.(string); ok && str == "red" {
				noRed = false
				break
			}
		}
		if noRed {
			for _, value := range inp {
				numbers = append(numbers, processInput(value)...)
			}
		}
	case float64:
		numbers = append(numbers, int(inp))
	}
	return numbers
}

func main(){
 f, _:= ioutil.ReadFile("/Users/aarayambeth/go_stuff/advent_of_code/cmd/JSAbacusFramework.io/input.txt")
 result := 0

 data := make(map[string]interface{},0)
 json.Unmarshal(f,&data)
 for _, num := range processInput(data) {
 	result += num
 }

 println(result)
}
