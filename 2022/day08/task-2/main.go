package main

import (
	"day08/trees"
	"fmt"
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

	grid, err := trees.GridFromInput(f)
	if err != nil {
		log.Fatalln(err)
	}

	sc, err := grid.HighestScenicScore()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(fmt.Sprintf("the highest scenic score is %d", sc))
}
