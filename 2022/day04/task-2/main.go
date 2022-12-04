package main

import (
	"bufio"
	"day04/cleaning"
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2022/day/4
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	overlappingPairs := 0
	for sc.Scan() {
		p, err := cleaning.PairFromString(sc.Text())
		if err != nil {
			log.Fatalln(err)
		}
		if p.OverlapEachOther() {
			overlappingPairs++
		}
	}
	log.Println(fmt.Sprintf("%d pairs overlap each other", overlappingPairs))
}
