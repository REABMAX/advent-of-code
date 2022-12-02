package main

import (
	"bufio"
	"day02/game"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/1
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	p := 0
	for sc.Scan() {
		l := sc.Text()
		pair := strings.Split(l, " ")
		if len(pair) < 2 {
			log.Fatalln("something is wrong with the input file's format")
		}

		opponentChoice := game.Shape(pair[0])
		myChoice := game.Shape(pair[1])

		p += myChoice.Fight(opponentChoice) + myChoice.Points()
	}

	log.Println(fmt.Sprintf("strategy guide sum: %d", p))
}
