package main

import (
	"../../pkg/file_reader"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getTotalArea(dimensionString string) (int,int) {
	var localSum int
	var localRibbon int
	s := strings.Split(dimensionString, "x")
	ints := make([]int, len(s))
	for i, str := range s {
		ints[i],_ = strconv.Atoi(str)
	}
	sort.Slice(ints, func(i int, j int) bool{
		return ints[i] < ints[j]
	})
	localSum += ints[1] * 3 * ints[0]
	localSum += ints[1] * ints[2] * 2
	localSum += ints[0] * ints[2] *2

	localRibbon += 2*(ints[0]+ints[1]) + ints[0]*ints[1]*ints[2]
	return localSum,localRibbon
}

func main(){
	fptr := flag.String("fpath","/Users/aarayambeth/go_stuff/advent_of_code/cmd/I_was_told_there_would_be_no_math/input.txt","path to input string")
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
	var totalWrapper int
	var totalRibbon int
	for s.Scan(){
		 x,y := getTotalArea(s.Text())
		 totalWrapper += x
		 totalRibbon += y
	}
	fmt.Printf("The required wrapping paper is %d square feet\nThe required ribbon is %d feet", totalWrapper, totalRibbon)
}
