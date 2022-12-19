package main

import (
	"day10/monkey"
	"log"
	"os"
	"sort"
)

// https://adventofcode.com/2022/day/8
func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	m, err := monkey.Parse(string(f))
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 20; i++ {
		monkey.PlayRound(true, m...)
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i].InspectCount() > m[j].InspectCount()
	})

	log.Println(m[0].InspectCount() * m[1].InspectCount())
}
