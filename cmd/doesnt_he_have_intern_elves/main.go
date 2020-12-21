package main

import (
	"../../pkg/file_reader"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func isNice(input string) int {
	reVowel := regexp.MustCompile(`[aeiou]`)
	//reRepeat := regexp.MustCompile(`([a-zA-Z])\\1`)
	//reExclude := regexp.MustCompile(`^(?!.*(ab|cd|pq|xy))`)
	reInclude := regexp.MustCompile(`^.*(?:(ab|cd|pq|xy)).*`)
	/*if !reInclude.MatchString(input) && len(reVowel.FindAllString(input,-1)) >= 3 && reRepeat.MatchString(input) {
		return 1
	}*/
	if !reInclude.MatchString(input){
		if len(reVowel.FindAllString(input,-1)) >= 3 {
			for i := 0; i < len(input)-1; i++ {
				if input[i+1] == input[i] {
					return 1
				}
			}
		} else {
			fmt.Println("Couldnt find 3 vowels")
		}
	} else {
		fmt.Println("reinclude check failed")
	}
	return 0
}

func isNice2(input string) int {
	var cond1, cond2 bool
	for i:=0;i<len(input)-2;i++ {
		if input[i]==input[i+2]{
			cond1 = true
			break
		}
	}

	for i:=0;i<len(input)-3;i++{
		for j:=i+2;j<len(input)-1;j++{
			if input[i]==input[j] && input[i+1]==input[j+1]{
				cond2 = true
				break
			}
		}
		if cond2==true{
			break
		}
	}
	if cond1 && cond2 {
		return 1
	}
	return 0
}

func main() {
	fptr := flag.String("fpath","/Users/aarayambeth/go_stuff/advent_of_code/cmd/doesnt_he_have_intern_elves/input.txt","enter path to input file")
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
	total := 0
	for s.Scan() {
		total += isNice2(s.Text())
	}
	fmt.Printf("%d strings are nice\n", total)
}
