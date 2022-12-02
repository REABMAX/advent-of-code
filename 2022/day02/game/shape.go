package game

import (
	"fmt"
	"log"
)

const (
	POINTS_LOSE = 0
	POINTS_DRAW = 3
	POINTS_WIN  = 6

	POINTS_STONE    = 1
	POINTS_PAPER    = 2
	POINTS_SCISSORS = 3
)

type Shape string

const STONE = Shape("A")
const PAPER = Shape("B")
const SCISSORS = Shape("C")

func (s Shape) Points() int {
	switch {
	case s.isStone():
		return POINTS_STONE
	case s.isPaper():
		return POINTS_PAPER
	case s.isScissors():
		return POINTS_SCISSORS
	}

	log.Fatalln(fmt.Sprintf("someone tried to cheat with an unsupported shape: %s", s))
	return 0
}
func (s Shape) Inferior() Shape {
	if s.isStone() {
		return SCISSORS
	}
	if s.isPaper() {
		return STONE
	}
	return PAPER
}
func (s Shape) Superior() Shape {
	if s.isStone() {
		return PAPER
	}
	if s.isPaper() {
		return SCISSORS
	}
	return STONE
}
func (s Shape) Fight(os Shape) int {
	if os == s.Superior() {
		return POINTS_LOSE
	}

	if os == s.Inferior() {
		return POINTS_WIN
	}

	return POINTS_DRAW
}

func (s Shape) isStone() bool    { return s == "A" || s == "X" }
func (s Shape) isPaper() bool    { return s == "B" || s == "Y" }
func (s Shape) isScissors() bool { return s == "C" || s == "Z" }
