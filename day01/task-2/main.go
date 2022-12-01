package main

import (
	"day01/calories"
	"fmt"
	"log"
)

// https://adventofcode.com/2022/day/1
func main() {
	cl, err := calories.ReadCaloriesList("calories.txt")
	if err != nil {
		log.Fatalln(err)
	}

	println("Solution 2")
	cl.SortByCalories()
	top3 := cl[0:3]
	sum := 0
	outList := ""
	for i, elf := range top3 {
		outList += fmt.Sprintf("%d) Elf number %d (%d calories)\n", i+1, elf.Number(), elf.Calories())
		sum += elf.Calories()
	}
	println(fmt.Sprintf("The top three elves with most calories in sum (%d) are:", sum))
	println(outList)
}
