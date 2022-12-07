package main

import (
	"day07/filesystem"
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

	dir := dirs.ChooseDirToFreeSpace(70_000_000, 30_000_000)
	log.Println(dir.Size())
}
