package calories

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	number   int
	calories int
}

func (e *elf) Number() int   { return e.number }
func (e *elf) Calories() int { return e.calories }

type CaloryList []elf

func ReadFromFile(fileName string) (CaloryList, error) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	i := 0
	cl := newCaloryList()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			i++
			cl.newElf()
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		cl[i].calories += cal
	}

	return cl, nil
}

func (c CaloryList) SortByCalories() {
	sort.Slice(c, func(i, j int) bool {
		return c[i].calories > c[j].calories
	})
}

func newCaloryList() CaloryList {
	return CaloryList{
		{
			number:   1,
			calories: 0,
		},
	}
}

func (c *CaloryList) newElf() {
	*c = append(*c, elf{
		number:   len(*c) + 1,
		calories: 0,
	})
}
