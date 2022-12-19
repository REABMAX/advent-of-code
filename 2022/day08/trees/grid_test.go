package trees

import (
	"fmt"
	"strings"
	"testing"
)

func TestGridFromInput(t *testing.T) {
	r := strings.NewReader(`30373
25512
65332
33549
35390`)
	exp := Grid{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	g, err := GridFromInput(r)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for i, v := range exp {
		for i2, v2 := range v {
			if g[i][i2] != v2 {
				t.Log(fmt.Sprintf("expected grid row %d column %d to be %d but it was %d instead", i, i2, v2, g[i][i2]))
				t.Fail()
			}
		}
	}
}

func TestGrid_VisibleTrees(t *testing.T) {
	g := Grid{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	exp := 21
	res, err := g.VisibleTrees()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if res != exp {
		t.Log(fmt.Sprintf("expected %d trees to be visible, was %d instead", exp, res))
		t.Fail()
	}
}

func TestGrid_ScenicScore(t *testing.T) {
	g := Grid{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	if sc, _ := g.ScenicScore(1, 2); sc != 4 {
		t.Log(fmt.Sprintf("Tree 1, 2 should have scenic score 4, was %d instead", sc))
		t.Fail()
	}

	if sc, _ := g.ScenicScore(3, 2); sc != 8 {
		t.Log(fmt.Sprintf("Tree 3, 2 should have scenic score 8, was %d instead", sc))
		t.Fail()
	}
}

func TestGrid_HighestScenicScore(t *testing.T) {
	g := Grid{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	h, _ := g.HighestScenicScore()
	if h != 8 {
		t.Log(fmt.Sprintf("The highest scenic score should be 8, was %d instead", h))
		t.Fail()
	}
}

func TestGrid_IsTreeVisible(t *testing.T) {
	g := Grid{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	if v, _ := g.IsTreeVisible(1, 1); !v {
		t.Log("Tree 1, 1 should be visible")
		t.Fail()
	}

	if v, _ := g.IsTreeVisible(1, 2); !v {
		t.Log("Tree 1, 2 should be visible")
		t.Fail()
	}

	if v, _ := g.IsTreeVisible(1, 3); v {
		t.Log("Tree 1, 3 should not be visible")
		t.Fail()
	}

	if v, _ := g.IsTreeVisible(2, 1); !v {
		t.Log("Tree 2, 1 should be visible")
		t.Fail()
	}

	if v, _ := g.IsTreeVisible(2, 2); v {
		t.Log("Tree 2, 2 should not be visible")
		t.Fail()
	}

	if v, _ := g.IsTreeVisible(4, 4); !v {
		t.Log("Tree 4, 4 should be visible")
		t.Fail()
	}
}
