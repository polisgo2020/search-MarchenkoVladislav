package main

import (
	"log"
	"os"
	"polisgomarchenko/index"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Invalid args number! Need 2 arguments: directory path and output file")
	}

	err := index.WriteInvertedIndexIntoFile(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}
