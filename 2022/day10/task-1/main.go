package main

import (
	"day10/signal"
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

	reg := signal.DefaultRegister

	err = reg.ExecuteProgram(f)
	if err != nil {
		log.Fatalln(err)
	}

	sum := 0
	sum += reg.SignalStrengthAtCycle(20)
	sum += reg.SignalStrengthAtCycle(60)
	sum += reg.SignalStrengthAtCycle(100)
	sum += reg.SignalStrengthAtCycle(140)
	sum += reg.SignalStrengthAtCycle(180)
	sum += reg.SignalStrengthAtCycle(220)

	log.Println(fmt.Sprintf("sum %d", sum))
}
