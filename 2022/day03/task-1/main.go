package main

import (
	"bufio"
	"day02/rucksack"
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2022/day/3
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	psum := 0
	for sc.Scan() {
		l := sc.Text()
		c := rucksack.CompartmentsFromString(l)
		psum += c.FindProblematicItemType().Priority()
	}

	log.Println(fmt.Sprintf("the priority sum of the problematic item types is %d", psum))
}
