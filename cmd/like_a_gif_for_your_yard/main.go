package main

import (
	"bufio"
	"os"
	"strings"
)

func animate(steps int, oldGrid [][]string) [][]string{
	var newGrid [][]string
	rowPositions := []int{-1,-1,-1,0,0,1,1,1}
	colPositions := []int{-1,0,1,-1,1,-1,0,1}
	for step:=0;step<steps;step++{

		newGrid = make([][]string, len(oldGrid))
		for i:=0;i<len(oldGrid);i++{
			newGrid[i] = make([]string, len(oldGrid[i]))
			copy(newGrid[i],oldGrid[i])
		}
		for i:=1;i<len(oldGrid)-1;i++{
			for j:=1;j<len(oldGrid[i])-1;j++{
				var on int
				for k:=0;k<8;k++{
					newRow := i+rowPositions[k]
					newColumn := j+colPositions[k]
					if oldGrid[newRow][newColumn] == "#" {
						on += 1
					}
				}
				if oldGrid[i][j] == "#" {
					// onlight tunrs off if on!=2 AND on!=3
					if on!=2 && on!=3 {
						newGrid[i][j] = "."
					}
				} else {
					// off light turns on if on==3
					if on==3{
						newGrid[i][j] = "#"
					}
				}
			}
		}
		newGrid[1][1] = "#"
		newGrid[1][len(newGrid[0])-2] = "#"
		newGrid[len(newGrid)-2][1] = "#"
		newGrid[len(newGrid)-2][len(newGrid[0])-2] = "#"
		oldGrid = newGrid
	}
	return newGrid
}

func main() {
	f, _ := os.Open("/Users/aarayambeth/go_stuff/advent_of_code/cmd/like_a_gif_for_your_yard/input.txt")
	s := bufio.NewScanner(f)
	var input [][]string
	//var i int
	input = make([][]string,0)
	for s.Scan(){
		chars := strings.Split(s.Text(),"")
		input = append(input, chars)
	}
	emptyRow := make([]string, len(input[0])+2)
	for i:=0;i<len(emptyRow);i++{
		emptyRow[i] = "."
	}
	for i:=0;i<len(input);i++{
		input[i] = append([]string{"."},input[i]...)
		input[i] = append(input[i],".")
	}
	input = append([][]string{emptyRow}, input...)
	input = append(input, emptyRow)
	result := animate(100,input)
	var total int
	for i:=1;i<len(result);i++{
		for j:=0;j<len(result[i]);j++{
			if result[i][j] == "#"{
				total += 1
			}
		}
	}
	println(total)
}
