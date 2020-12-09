package file_reader

import (
	"log"
	"os" )

func GetFileHandle(path *string)  *os.File {
	f,err := os.Open(*path)
	if err!=nil{
		log.Fatal(err)
		return nil
	}
	return f
}
