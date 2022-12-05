package rearrangement

import (
	"fmt"
	"testing"
)

func TestParseStacks(t *testing.T) {
	dia := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `
	stacks := ParseStacks(dia)
	exp := Stacks{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}
	assertExpectedStacks(t, stacks, exp)
}

func TestStacks_Move(t *testing.T) {
	stacks := Stacks{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}
	exp := Stacks{
		{'Z', 'N'},
		{'M'},
		{'P', 'D', 'C'},
	}
	err := stacks.Move(2, 2, 3)
	if err != nil {
		t.Log(fmt.Sprintf("didnt expect error %v", err))
		t.Fail()
	}

	assertExpectedStacks(t, stacks, exp)
}

func TestStacks_MoveMultipleAtOnce(t *testing.T) {
	stacks := Stacks{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}
	exp := Stacks{
		{'Z', 'N'},
		{'M'},
		{'P', 'C', 'D'},
	}
	err := stacks.MoveMultipleAtOnce(2, 2, 3)
	if err != nil {
		t.Log(fmt.Sprintf("didnt expect error %v", err))
		t.Fail()
	}

	assertExpectedStacks(t, stacks, exp)
}

func TestParseMovement(t *testing.T) {
	s := "move 2 from 1 to 3"
	q, from, to, err := ParseMovement(s)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if q != 2 {
		t.Log(fmt.Sprintf("expected quantity to be 2, is %d instead", q))
		t.Fail()
	}
	if from != 1 {
		t.Log(fmt.Sprintf("expected from to be 1, is %d instead", from))
		t.Fail()
	}
	if to != 3 {
		t.Log(fmt.Sprintf("expected to to be 3, is %d instead", to))
		t.Fail()
	}
}

func assertExpectedStacks(t *testing.T, stacks Stacks, exp Stacks) {
	for nr, s := range stacks {
		for i, c := range s {
			e := exp[nr][i]
			if c != e {
				t.Log(fmt.Sprintf("expected stack %d crate %d to be %v but it was %v instead", nr, i, e, c))
				t.Fail()
			}
		}
	}
}
