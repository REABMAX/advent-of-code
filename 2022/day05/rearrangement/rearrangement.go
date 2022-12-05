package rearrangement

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type crate rune

func (c crate) isValid() bool {
	return c >= 65 && c <= 90
}

type stack []crate

func (s *stack) prepend(c crate) {
	*s = append([]crate{c}, *s...)
}

func (s *stack) pop() (crate, error) {
	if len(*s) <= 0 {
		return 0, errors.New("could not pop from stack since it's empty")
	}
	sv := *s
	i := len(sv) - 1
	c := sv[i]
	*s = append(sv[:i], stack{}...)
	return c, nil
}

type Stacks []stack

func (s Stacks) Move(quantity, from, to int) error {
	from -= 1
	to -= 1

	if !s.hasStackNr(from) {
		return errors.New(fmt.Sprintf("stack nr %d does not exist", from))
	}
	if !s.hasStackNr(to) {
		return errors.New(fmt.Sprintf("stack nr %d does not exist", to))
	}
	if len(s[from]) < quantity {
		return errors.New(fmt.Sprintf("stack nr %d does not contain enough crates", from))
	}

	for i := 0; i < quantity; i++ {
		c, err := s[from].pop()
		if err != nil {
			return err
		}
		s[to] = append(s[to], c)
	}

	return nil
}

func (s Stacks) MoveMultipleAtOnce(quantity, from, to int) error {
	from -= 1
	to -= 1

	if !s.hasStackNr(from) {
		return errors.New(fmt.Sprintf("stack nr %d does not exist", from))
	}
	if !s.hasStackNr(to) {
		return errors.New(fmt.Sprintf("stack nr %d does not exist", to))
	}
	if len(s[from]) < quantity {
		return errors.New(fmt.Sprintf("stack nr %d does not contain enough crates", from))
	}

	var crates []crate
	for i := 0; i < quantity; i++ {
		c, err := s[from].pop()
		if err != nil {
			return err
		}
		crates = append([]crate{c}, crates...)
	}
	s[to] = append(s[to], crates...)

	return nil
}

func (s Stacks) TopCrates() []string {
	var crates []string
	for _, st := range s {
		crates = append(crates, string(st[len(st)-1]))
	}
	return crates
}

func (s Stacks) hasStackNr(nr int) bool {
	return len(s) > nr
}

func ParseStacks(dia string) Stacks {
	sc := bufio.NewScanner(strings.NewReader(dia))
	stacks := Stacks{}
	for sc.Scan() {
		runes := []rune(sc.Text())
		if len(runes) <= 0 {
			break
		}
		stackNr := 0
		for i := 1; i < len(runes); i += 4 {
			if !stacks.hasStackNr(stackNr) {
				stacks = append(stacks, stack{})
			}
			c := crate(runes[i])
			if c.isValid() {
				stacks[stackNr].prepend(c)
			}
			stackNr++
		}
	}

	return stacks
}

func ParseMovement(s string) (int, int, int, error) {
	reg := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	m := reg.FindStringSubmatch(s)
	if len(m) < 4 {
		return 0, 0, 0, errors.New("move action string: wrong format")
	}

	q, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, 0, 0, err
	}
	from, err := strconv.Atoi(m[2])
	if err != nil {
		return 0, 0, 0, err
	}
	to, err := strconv.Atoi(m[3])
	if err != nil {
		return 0, 0, 0, err
	}

	return q, from, to, nil
}
