package main

import (
	"../../pkg/file_reader"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)


func getFloor(address string) (bool,int) {
	var sum int
	for pos,runeValue := range address{
		if runeValue == '(' {
			sum += 1
		} else if runeValue == ')'{
			sum -=1
		}
		if sum < 0 {
			return true,(pos+1)
		}
	}
	return false,sum
}

func main() {
	fptr := flag.String("fpath","","file path to read from")
	flag.Parse()
	f := file_reader.GetFileHandle(fptr)
	if f == nil {
		log.Fatal("Could not read file: Could not get File Handle")
		os.Exit(-1)
	}
	defer func(){
		if err := f.Close(); err!=nil{
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan(){
		isBasement, val := getFloor(s.Text())
		if isBasement {
			fmt.Printf("Entered the basement at position %d\n", val)
		} else {
			fmt.Printf("The floor is :%d\n", val)
		}
	}
}