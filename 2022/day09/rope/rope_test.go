package rope

import (
	"fmt"
	"testing"
)

func TestGrid_Simulate(t *testing.T) {
	cmds := []Command{
		{"R", 4},
		{"U", 4},
		{"L", 3},
		{"D", 1},
		{"R", 4},
		{"D", 1},
		{"L", 5},
		{"R", 2},
	}
	g := DefaultRope
	g.Simulate(cmds)

	if len(g.backbrain) != 13 {
		t.Log(fmt.Sprintf("expected backbrain to have length of 13 but was %d instead", len(g.backbrain)))
		t.Fail()
	}
}

func TestGrid_Simulate_LongRope(t *testing.T) {
	cmds := []Command{
		{"R", 5},
		{"U", 8},
		{"L", 8},
		{"D", 3},
		{"R", 17},
		{"D", 10},
		{"L", 25},
		{"U", 20},
	}
	g := NewRope(10)
	g.Simulate(cmds)

	if len(g.backbrain) != 36 {
		t.Log(fmt.Sprintf("expected backbrain to have length of 36 but was %d instead", len(g.backbrain)))
		t.Log(g.backbrain)
		t.Fail()
	}
}

func TestGrid_Up(t *testing.T) {
	g := DefaultRope

	g.Up(5)

	if g.head.Row() != 5 {
		t.Log(fmt.Sprintf("head row should be 5, is %d", g.head.Row()))
		t.Fail()
	}
	if g.tail.Row() != 4 {
		t.Log(fmt.Sprintf("tail row should be 4, is %d", g.tail.Row()))
		t.Fail()
	}
}

func TestGrid_Down(t *testing.T) {
	g := DefaultRope

	g.Down(5)

	if g.head.Row() != -5 {
		t.Log(fmt.Sprintf("head row should be -5, is %d", g.head.Row()))
		t.Fail()
	}
	if g.tail.Row() != -4 {
		t.Log(fmt.Sprintf("tail row should be -4, is %d", g.tail.Row()))
		t.Fail()
	}
}

func TestGrid_Left(t *testing.T) {
	g := DefaultRope

	g.Left(5)

	if g.head.Col() != -5 {
		t.Log(fmt.Sprintf("head col should be -5, is %d", g.head.Col()))
		t.Fail()
	}
	if g.tail.Col() != -4 {
		t.Log(fmt.Sprintf("tail col should be -4, is %d", g.tail.Col()))
		t.Fail()
	}
}

func TestGrid_Right(t *testing.T) {
	g := DefaultRope

	g.Right(5)

	if g.head.Col() != 5 {
		t.Log(fmt.Sprintf("head col should be 5, is %d", g.head.Col()))
		t.Fail()
	}
	if g.tail.Col() != 4 {
		t.Log(fmt.Sprintf("tail col should be 4, is %d", g.tail.Col()))
		t.Fail()
	}
}
