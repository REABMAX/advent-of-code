package main

import (
	"day10/signal"
	"log"
	"os"
)

// https://adventofcode.com/2022/day/8
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	reg := signal.DefaultRegister

	err = reg.ExecuteProgram(f)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("\n" + reg.DrawCRT())
}
