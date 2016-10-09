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
	// // generate 0-23h as float
	// var hours []float32
	// hours = make([]float32, 24)
	// for i := 0; i < 24; i++ {
	// 	hours[i] = float32(i) / float32(24)
	// 	fmt.Printf("hour[%d]: %v\n", i, hours[i])
	// }

	//process
	err := process(file)
	if err != nil {
		log.Printf("process file error: %v\n", err)
		return
	}

	//write
}
