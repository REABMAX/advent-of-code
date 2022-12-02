package main

import (
	"day01/calories"
	"fmt"
	"log"
)

// https://adventofcode.com/2022/day/1
func main() {
	cl, err := calories.ReadFromFile("calories.txt")
	if err != nil {
		log.Fatalln(err)
	}

	println("Solution 1")
	cl.SortByCalories()
	elf := cl[0]
	println(fmt.Sprintf("The elf with most calories (%d) is elf number %d.", elf.Calories(), elf.Number()))
}
