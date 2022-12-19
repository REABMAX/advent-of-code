package trees

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type Grid [][]tree

type tree int

func (t tree) isVisibleAmong(ts []tree) bool {
	for _, ot := range ts {
		if ot >= t {
			return false
		}
	}
	return true
}

// Calculate viewingDistance while the given []tree order has to be: the first in the row is the one next to t
func (t tree) viewingDistance(ts []tree) int {
	dist := 0
	for ; dist < len(ts); dist++ {
		if ts[dist] >= t {
			dist++
			break
		}
	}

	if dist > 0 {
		return dist
	}

	return 1
}

func GridFromInput(r io.Reader) (Grid, error) {
	sc := bufio.NewScanner(r)

	var g Grid
	for sc.Scan() {
		line := sc.Text()
		var row []tree
		for _, char := range line {
			t, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			row = append(row, tree(t))
		}
		g = append(g, row)
	}

	return g, nil
}

func (g Grid) VisibleTrees() (int, error) {
	c := 0
	for ri, r := range g {
		for ci, _ := range r {
			v, err := g.IsTreeVisible(ri, ci)
			if err != nil {
				return 0, err
			}
			if v {
				c++
			}
		}
	}
	return c, nil
}

func (g Grid) HighestScenicScore() (int, error) {
	highest := 0
	for ri, r := range g {
		for ci, _ := range r {
			sc, err := g.ScenicScore(ri, ci)
			if err != nil {
				return 0, err
			}
			if sc > highest {
				highest = sc
			}
		}
	}
	return highest, nil
}

func (g Grid) IsTreeVisible(rowNr, colNr int) (bool, error) {
	if len(g) <= rowNr {
		return false, errors.New("invalid coordinates: row is not available")
	}

	row := g[rowNr]
	if len(row) <= colNr {
		return false, errors.New("invalid coordinates: col is not available")
	}

	t := row[colNr]
	top, right, bottom, left := getTreesInTheWay(g, rowNr, colNr)

	visible := t.isVisibleAmong(top) || t.isVisibleAmong(right) || t.isVisibleAmong(bottom) || t.isVisibleAmong(left)
	return visible, nil
}

func (g Grid) ScenicScore(rowNr, colNr int) (int, error) {
	if len(g) <= rowNr {
		return 0, errors.New("invalid coordinates: row is not available")
	}

	row := g[rowNr]
	if len(row) <= colNr {
		return 0, errors.New("invalid coordinates: col is not available")
	}

	t := row[colNr]
	top, right, bottom, left := getTreesInTheWay(g, rowNr, colNr)

	return t.viewingDistance(reverse(top)) * t.viewingDistance(right) * t.viewingDistance(bottom) * t.viewingDistance(reverse(left)), nil
}

func getTreesInTheWay(g Grid, rowNr, colNr int) ([]tree, []tree, []tree, []tree) {
	var top []tree
	var right []tree
	var bottom []tree
	var left []tree
	for ri, r := range g {
		if ri < rowNr {
			top = append(top, r[colNr])
		}
		if ri == rowNr && colNr > 0 && colNr < len(g[rowNr])-1 {
			for ci, c := range r {
				if ci < colNr {
					left = append(left, c)
				}
				if ci > colNr {
					right = append(right, c)
				}
			}
		}
		if ri > rowNr {
			bottom = append(bottom, r[colNr])
		}
	}
	return top, right, bottom, left
}

func reverse[T comparable](input []T) []T {
	il := len(input)
	out := make([]T, il)
	for i, n := range input {
		j := il - i - 1
		out[j] = n
	}
	return out
}
