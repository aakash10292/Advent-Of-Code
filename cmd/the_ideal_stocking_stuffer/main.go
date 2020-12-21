package main

import (
	"../../pkg/file_reader"
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func solve(input string) int {
	re := regexp.MustCompile("^000000")
	var result int
	msg := ""
	for !re.MatchString(msg){
		result++
		hash := md5.New()
		io.WriteString(hash, input)
		io.WriteString(hash, fmt.Sprintf("%d",result))
		msg = fmt.Sprintf("%x",hash.Sum(nil))
	}
	return result
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
		fmt.Printf("The answer is %d\n", solve(s.Text()))
	}
}