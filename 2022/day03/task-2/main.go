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
	lines := []string{}
	for sc.Scan() {
		lines = append(lines, sc.Text())
		if len(lines) == 3 {
			g := rucksack.GroupFromStrings(lines...)
			psum += g.FindUnifyingItemType().Priority()
			lines = []string{}
		}
	}

	log.Println(fmt.Sprintf("the priority sum of the elf groups is %d", psum))
}
