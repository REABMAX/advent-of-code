package main

import (
	"bufio"
	"day09/rope"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/8
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var cmds []rope.Command
	for sc.Scan() {
		line := sc.Text()
		cmd := strings.Split(line, " ")
		steps, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatalln(err)
		}
		cmds = append(cmds, rope.NewCommand(cmd[0], steps))
	}

	r := rope.DefaultRope
	r.Simulate(cmds)

	log.Println(fmt.Sprintf("the tail visited %d positions at least once", r.TailPositions()))
}
