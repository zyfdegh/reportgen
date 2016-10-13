package main

import (
	"log"
	"os"
)

func main() {
	// open
	file := "C:\\demo.xls"
	if _, err := os.Stat(file); err != nil {
		log.Printf("stat file error: %v\n", err)
		return
	}

	//process
	err := process(file)
	if err != nil {
		log.Printf("process file error: %v\n", err)
		return
	}

	//write
}
