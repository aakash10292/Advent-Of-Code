package main

import (
	"../../pkg/file_reader"
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)




func main() {
	fptr := flag.String("fpath", "/Users/aarayambeth/go_stuff/advent_of_code/cmd/probably_a_fire_hazard/input.txt", "enter path to input file")
	flag.Parse()
	f := file_reader.GetFileHandle(fptr)
	if f == nil {
		log.Fatal("Could not open file : Failed to get File handle")
		os.Exit(-1)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	var grid [1000][1000]int

	fromRegex := regexp.MustCompile("\\d+,\\d+")
	toRegex := regexp.MustCompile("\\d+,\\d+$")

	for s.Scan() {
		instruction := s.Text()

		fromStr := fromRegex.FindString(instruction)
		toStr := toRegex.FindString(instruction)
		from := strings.Split(fromStr, ",")
		to := strings.Split(toStr, ",")

		fromX, _ := strconv.Atoi(from[0])
		fromY, _ := strconv.Atoi(from[1])
		toX, _ := strconv.Atoi(to[0])
		toY, _ := strconv.Atoi(to[1])

		for x := fromX; x <= toX; x++ {
			for y := fromY; y <= toY; y++ {
				if strings.Contains(instruction, "off") {
					if grid[x][y]>0{
						grid[x][y] -= 1
					}
				} else if strings.Contains(instruction, "on") {
					grid[x][y] += 1
				} else {
					grid[x][y] += 2
				}
			}
		}
	}

	total := 0

	for _, row := range grid {
		for _, col := range row {
			total += col
		}
	}

	println(total)
}
