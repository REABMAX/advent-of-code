package main

import (
	"bufio"
	"day05/rearrangement"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/5
func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	parts := strings.Split(string(f), "\n\n")
	stacks := rearrangement.ParseStacks(parts[0])

	sc := bufio.NewScanner(strings.NewReader(parts[1]))
	for sc.Scan() {
		l := sc.Text()
		quantity, from, to, err := rearrangement.ParseMovement(l)
		if err != nil {
			log.Fatalln(err)
		}

		err = stacks.MoveMultipleAtOnce(quantity, from, to)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println(fmt.Sprintf("Top crates message: %s", strings.Join(stacks.TopCrates(), "")))
}
