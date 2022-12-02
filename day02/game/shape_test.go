package game

import (
	"fmt"
	"testing"
)

func TestShapeMapsCorrectly(t *testing.T) {
	shape := Shape("A")
	if !shape.isStone() {
		t.Fail()
	}
	shape = "X"
	if !shape.isStone() {
		t.Fail()
	}
	shape = "B"
	if !shape.isPaper() {
		t.Fail()
	}
	shape = "Y"
	if !shape.isPaper() {
		t.Fail()
	}
	shape = "C"
	if !shape.isScissors() {
		t.Fail()
	}
	shape = "Z"
	if !shape.isScissors() {
		t.Fail()
	}
}

func TestShapeReturnsCorrectShapePoints(t *testing.T) {
	shape := Shape("A")
	if shape.Points() != 1 {
		t.Log(fmt.Sprintf("stone must give 1 point exactly, %d points where given", shape.Points()))
		t.Fail()
	}
	shape = "B"
	if shape.Points() != 2 {
		t.Log(fmt.Sprintf("paper must give 2 point exactly, %d points where given", shape.Points()))
		t.Fail()
	}
	shape = "C"
	if shape.Points() != 3 {
		t.Log(fmt.Sprintf("stone must give 3 point exactly, %d points where given", shape.Points()))
		t.Fail()
	}
}

func TestShapeReturnsCorrectInferior(t *testing.T) {
	stone := Shape("A")
	paper := Shape("B")
	scissors := Shape("C")

	if stone.Inferior() != scissors {
		t.Log(fmt.Sprintf("stone's inferior is scissors but %s was returned", stone.Inferior()))
		t.Fail()
	}

	if paper.Inferior() != stone {
		t.Log(fmt.Sprintf("paper's inferior is stone but %s was returned", paper.Inferior()))
		t.Fail()
	}

	if scissors.Inferior() != paper {
		t.Log(fmt.Sprintf("scissor's inferior is paper but %s was returned", scissors.Inferior()))
		t.Fail()
	}
}

func TestShapeReturnsCorrectSuperior(t *testing.T) {
	stone := Shape("A")
	paper := Shape("B")
	scissors := Shape("C")

	if stone.Superior() != paper {
		t.Log(fmt.Sprintf("stone's superior is paper but %s was returned", stone.Superior()))
		t.Fail()
	}

	if paper.Superior() != scissors {
		t.Log(fmt.Sprintf("paper's superior is scissors but %s was returned", paper.Superior()))
		t.Fail()
	}

	if scissors.Superior() != stone {
		t.Log(fmt.Sprintf("scissor's inferior is stone but %s was returned", scissors.Superior()))
		t.Fail()
	}
}

func TestShapeFightReturnsCorrectPoints(t *testing.T) {
	stone := Shape("A")
	paper := Shape("B")
	scissors := Shape("C")

	// stone
	if stone.Fight(paper) != POINTS_LOSE {
		t.Log(fmt.Sprintf("stone should lose against paper but %d points were given instead", stone.Fight(paper)))
		t.Fail()
	}

	if stone.Fight(scissors) != POINTS_WIN {
		t.Log(fmt.Sprintf("stone should win against scissors but %d points were given instead", stone.Fight(scissors)))
		t.Fail()
	}

	if stone.Fight(stone) != POINTS_DRAW {
		t.Log(fmt.Sprintf("stone should draw against stone but %d points were given instead", stone.Fight(stone)))
		t.Fail()
	}

	// paper
	if paper.Fight(paper) != POINTS_DRAW {
		t.Log(fmt.Sprintf("paper should draw against paper but %d points were given instead", paper.Fight(paper)))
		t.Fail()
	}

	if paper.Fight(scissors) != POINTS_LOSE {
		t.Log(fmt.Sprintf("paper should lose against scissors but %d points were given instead", paper.Fight(scissors)))
		t.Fail()
	}

	if paper.Fight(stone) != POINTS_WIN {
		t.Log(fmt.Sprintf("paper should win against stone but %d points were given instead", paper.Fight(stone)))
		t.Fail()
	}

	// scissors
	if scissors.Fight(paper) != POINTS_WIN {
		t.Log(fmt.Sprintf("stone should win against paper but %d points were given instead", scissors.Fight(paper)))
		t.Fail()
	}

	if scissors.Fight(scissors) != POINTS_DRAW {
		t.Log(fmt.Sprintf("scissors should draw against scissors but %d points were given instead", scissors.Fight(scissors)))
		t.Fail()
	}

	if scissors.Fight(stone) != POINTS_LOSE {
		t.Log(fmt.Sprintf("scissors should lose against stone but %d points were given instead", scissors.Fight(stone)))
		t.Fail()
	}
}
