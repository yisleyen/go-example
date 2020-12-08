package main

import (
	"log"
	"os"
)

var filePath = "test.txt"

func main() {
	fileCreate()
}

func fileCreate() {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	defer f.Close()
}
