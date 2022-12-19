package rope

import (
	"fmt"
	"log"
)

type Command struct {
	direction string
	steps     int
}

func NewCommand(dir string, steps int) Command {
	return Command{dir, steps}
}

type coordinate [2]int

func (c coordinate) Row() int               { return c[0] }
func (c coordinate) Col() int               { return c[1] }
func (c coordinate) PositionString() string { return fmt.Sprintf("%d.%d", c.Row(), c.Col()) }

type Rope struct {
	length    int
	tail      coordinate
	head      coordinate
	backbrain map[string]struct{}
}

var DefaultRope = NewRope(2)

func NewRope(length int) Rope {
	return Rope{
		length:    length,
		tail:      coordinate{0, 0},
		head:      coordinate{0, 0},
		backbrain: make(map[string]struct{}),
	}
}

func (r *Rope) Simulate(cmds []Command) {
	for _, cmd := range cmds {
		switch cmd.direction {
		case "U":
			r.Up(cmd.steps)
		case "R":
			r.Right(cmd.steps)
		case "D":
			r.Down(cmd.steps)
		case "L":
			r.Left(cmd.steps)
		default:
			log.Fatalln("unsupported command")
		}
	}
}

func (r *Rope) TailPositions() int {
	return len(r.backbrain)
}

func (r *Rope) Up(n int) {
	for i := 0; i < n; i++ {
		r.head[0]++
		r.moveTail()
	}
}

func (r *Rope) Right(n int) {
	for i := 0; i < n; i++ {
		r.head[1]++
		r.moveTail()
	}
}

func (r *Rope) Down(n int) {
	for i := 0; i < n; i++ {
		r.head[0]--
		r.moveTail()
	}
}

func (r *Rope) Left(n int) {
	for i := 0; i < n; i++ {
		r.head[1]--
		r.moveTail()
	}
}

func (r *Rope) moveTail() {
	if diff(r.head.Row(), r.tail.Row()) >= r.length {
		if r.head.Row() > r.tail.Row() {
			r.tail[0]++
		} else {
			r.tail[0]--
		}
		r.tail[1] = r.head.Col()
	} else if diff(r.head.Col(), r.tail.Col()) >= r.length {
		if r.head.Col() > r.tail.Col() {
			r.tail[1]++
		} else {
			r.tail[1]--
		}
		r.tail[0] = r.head.Row()
	}
	r.backbrain[r.tail.PositionString()] = struct{}{}
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
