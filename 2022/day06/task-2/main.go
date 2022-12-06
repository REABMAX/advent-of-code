package main

import (
	"day06/signal"
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2022/day/6
func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	som, err := signal.GetStartOfMessage(string(f))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(fmt.Sprintf("start-of-message marker: %d", som))
}
