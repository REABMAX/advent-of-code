package main

import (
	"day07/filesystem"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/6
func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	ls := strings.Split(string(f), "\n")
	dirs, err := filesystem.DirsFromSlice(ls)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(fmt.Sprintf("The sum is %d", dirs.Dirs(func(d filesystem.Dir) bool {
		return d.Size() <= 100_000
	}).Size()))
}
