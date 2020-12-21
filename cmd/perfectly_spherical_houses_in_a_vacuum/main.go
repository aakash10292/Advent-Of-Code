package main

import (
	"../../pkg/file_reader"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type point struct {
	x int
	y int
}


func getHouses(input string) int {
	total := 1
	var curPos = make([]point,2)
	curPos[0] = point{0,0}
	curPos[1] = point{0,0}
	var santa = make(map[point]struct{})
	var roboSanta = make(map[point]struct{})
	var sets = make([]map[point]struct{},2)
	sets[0] = santa
	sets[1] = roboSanta
 	var member struct{}
	sets[0][curPos[0]] = member
	sets[1][curPos[1]] = member
	for pos ,symbol:= range input {
		if symbol == '^' {
			curPos[pos%2].y++
		} else if symbol == '>' {
			curPos[pos%2].x++
		} else if symbol == '<' {
			curPos[pos%2].x--
		} else if symbol == 'v' {
			curPos[pos%2].y--
		}

		_,exists1 := sets[0][curPos[pos%2]]
		_,exists2 := sets[1][curPos[pos%2]]
		if (exists1 || exists2)  == false {
			newPoint := point{curPos[pos%2].x,curPos[pos%2].y}
			sets[pos%2][newPoint] = member
			total++
		}
	}
	return total
}

func main(){
	fptr := flag.String("fpath","","enter path to input file")
	flag.Parse()
	f := file_reader.GetFileHandle(fptr)
	if f == nil {
		log.Fatal("Could not open file : Failed to get File handle")
		os.Exit(-1)
	}
	defer func(){
		if err:=f.Close(); err != nil{
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan(){
		fmt.Printf("%d houses got atleast one present\n", getHouses(s.Text()))
	}
}